package ocher

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"mkuznets.com/go/ocher/internal/pb"
)

const MaxRetries = 3

var IdentRe = regexp.MustCompile(`[^a-zA-Z0-9_]+`)

type Server struct {
	DB    *pgxpool.Pool
	grpc  *grpc.Server
	hooks map[string]*Hook
}

type Init struct {
	ClientID string
	Queue    string
}

type Hook struct {
	Queue string
	Pre   func(tx pgx.Tx, args []byte) error
	Post  func(tx pgx.Tx, args []byte, result []byte) error
}

func New(dsn string) *Server {

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	server := &Server{
		DB:    pool,
		hooks: make(map[string]*Hook),
		grpc:  grpc.NewServer(grpc.ConnectionTimeout(10 * time.Second)),
	}
	pb.RegisterOcherServer(server.grpc, server)

	return server
}

func (server *Server) AddHook(h *Hook) {
	server.hooks[h.Queue] = h
}

func (server *Server) Serve(lis net.Listener) error {
	return server.grpc.Serve(lis)
}

func (server *Server) QueueTransactional(s pb.Ocher_QueueTransactionalServer) error {

	init, err := server.waitForInit(s)
	if err != nil {
		log.Printf("init error: %v", err)
		return err
	}
	log.Printf("`%s' connected (queue: %s), listening to new tasks", init.ClientID, init.Queue)

	for foundTask := false; ; {

		err = Tx(s.Context(), server.DB, func(tx pgx.Tx) error {
			task := new(pb.Task)

			err = tx.QueryRow(s.Context(), `
			UPDATE "ocher_tasks" SET status='EXECUTING', started_at=STATEMENT_TIMESTAMP()
			WHERE id=(
				SELECT id FROM "ocher_tasks"
				WHERE status='ENQUEUED' AND queue=$1
				ORDER BY id
				FOR UPDATE SKIP LOCKED
				LIMIT 1
			) RETURNING id, queue, args
			;`, init.Queue).Scan(&task.Id, &task.Queue, &task.Args)

			switch {
			case err == pgx.ErrNoRows:
				return nil
			case err != nil:
				return errors.WithStack(err)
			}

			log.Printf("[INFO] got task: %v", task.Id)
			foundTask = true

			var retries int
			err = tx.QueryRow(s.Context(), `
			SELECT COUNT(*) FROM "ocher_retries" WHERE task_id=$1
			`, task.Id).Scan(&retries)

			if retries > MaxRetries {
				_, err = tx.Exec(s.Context(), `
				UPDATE "ocher_tasks" SET status='EXPIRED', finished_at=STATEMENT_TIMESTAMP() WHERE id=$1
				`, task.Id)
				if err != nil {
					return errors.WithStack(err)
				}
				return nil
			}

			if err := server.markRetry(s.Context(), task.Id); err != nil {
				return err
			}

			hook := server.hooks[task.Queue]

			if hook != nil && hook.Pre != nil {
				if err := hook.Pre(tx, task.Args); err != nil {
					return err
				}
			}

			// Send task
			if err := s.Send(task); err != nil {
				return err
			}

			// Get task status
			st, err := s.Recv()

			switch {
			case err == io.EOF:
				return err
			case err != nil:
				return errors.WithStack(err)
			}

			switch x := st.Event.(type) {

			case *pb.Status_Error_:
				_, err = tx.Exec(s.Context(), `
				UPDATE "ocher_tasks" SET status='FAILURE', result=$1, message=$2, finished_at=STATEMENT_TIMESTAMP() WHERE id=$3
				`, x.Error.Error, x.Error.Message, task.Id)
				if err != nil {
					return errors.WithStack(err)
				}

			case *pb.Status_Finish_:
				_, err = tx.Exec(s.Context(), `
				UPDATE "ocher_tasks" SET status='FINISHED', result=$1, finished_at=STATEMENT_TIMESTAMP() WHERE id=$2
				`, x.Finish.Result, task.Id)
				if err != nil {
					return errors.WithStack(err)
				}

				if hook != nil && hook.Post != nil {
					if err := hook.Post(tx, task.Args, x.Finish.Result); err != nil {
						_, err = tx.Exec(s.Context(), `
						UPDATE "ocher_tasks" SET status='FAILURE', message=$1, finished_at=STATEMENT_TIMESTAMP() WHERE id=$2
						`, err.Error(), task.Id)
						if err != nil {
							return errors.WithStack(err)
						}
					}
				}

			default:
				return errors.New("invalid status message")
			}

			return nil
		})
		if err != nil {
			log.Printf("%v", err)
			return err
		}

		if foundTask {
			foundTask = false
			continue
		} else {
			if err := server.waitForTask(s.Context(), init.Queue); err != nil {
				log.Printf("waitForTask error: %v", err)
				return err
			}
		}
	}
}

func (server *Server) waitForTask(ctx context.Context, queue string) error {
	conn, err := server.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	channelName := IdentRe.ReplaceAllLiteralString(queue, "_")

	if _, err := conn.Exec(ctx, fmt.Sprintf("LISTEN ocher_%s;", channelName)); err != nil {
		return err
	}

	waitCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, _ = conn.Conn().WaitForNotification(waitCtx)
	if _, err := conn.Exec(ctx, fmt.Sprintf("UNLISTEN ocher_%s;", channelName)); err != nil {
		return err
	}

	return nil
}

func (server *Server) markRetry(ctx context.Context, taskId uint64) error {
	return Tx(ctx, server.DB, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, `INSERT INTO "ocher_retries" (task_id) VALUES ($1)`, taskId)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

func (server *Server) waitForInit(s pb.Ocher_QueueTransactionalServer) (*Init, error) {
	in, err := s.Recv()

	switch {
	case err == io.EOF:
		return nil, nil
	case err != nil:
		return nil, err
	}

	msg, ok := in.Event.(*pb.Status_Init_)
	if !ok {
		return nil, errors.New("unexpected initial message")
	}

	if msg.Init.Id == "" {
		return nil, errors.New("init.id is not defined")
	}
	if msg.Init.Queue == "" {
		return nil, errors.New("init.queue is not defined")
	}

	init := &Init{
		ClientID: msg.Init.Id,
		Queue:    msg.Init.Queue,
	}

	return init, nil
}
