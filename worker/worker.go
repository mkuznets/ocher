package worker

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"mkuznets.com/go/ocher/internal/pb"
)

const (
	dialTimeout       = 30
	reconnectAttempts = 3
)

type Task struct {
	ID    uint64
	Queue string
	Args  []byte
}

type Handler func(context.Context, *Task) ([]byte, error)

type Worker struct {
	target   string
	clientID string
	handlers map[string]Handler
}

func New(target string, clientID string) *Worker {
	return &Worker{
		target:   target,
		clientID: clientID,
		handlers: make(map[string]Handler),
	}
}

func (q *Worker) TaskFunc(queue string, hx Handler) {
	q.handlers[queue] = hx
}

func (q *Worker) serveQueue(ctx context.Context, client pb.OcherClient, queue string) (error, bool) {
	hx := q.handlers[queue]

	streamCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := client.QueueTransactional(streamCtx, grpc.WaitForReady(true))
	if err != nil {
		return err, false
	}

	log.Printf("%s: connected", queue)

	msg := &pb.Status{
		Event: &pb.Status_Init_{
			Init: &pb.Status_Init{
				Id:    q.clientID,
				Queue: queue,
			},
		},
	}
	if err := stream.Send(msg); err != nil {
		return err, false
	}

	log.Printf("%s: initialised, waiting for tasks", queue)

	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil, true
			}
			return err, true
		}

		task := &Task{
			ID:    in.Id,
			Queue: in.Queue,
			Args:  in.Args,
		}

		r, err := hx(ctx, task)
		if err != nil {
			msg := &pb.Status{Event: &pb.Status_Error_{Error: &pb.Status_Error{Message: err.Error(), Error: r}}}
			if err := stream.Send(msg); err != nil {
				return err, true
			}
		} else {
			msg := &pb.Status{Event: &pb.Status_Finish_{Finish: &pb.Status_Finish{Result: r}}}
			if err := stream.Send(msg); err != nil {
				return err, true
			}
		}
	}
}

func (q *Worker) Serve() error {
	ctx := context.Background()

	dialCtx, cancel := context.WithTimeout(ctx, dialTimeout*time.Second)
	defer cancel()

	log.Printf("%s: dialing...", q.target)

	conn, err := grpc.DialContext(dialCtx, q.target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	log.Printf("%s: connected", q.target)

	defer func() {
		if e := conn.Close(); e != nil {
			err = e
			return
		}
	}()

	client := pb.NewOcherClient(conn)

	var wg sync.WaitGroup

	for queue := range q.handlers {
		wg.Add(1)

		go func(queue string) {
			for i := 0; i < reconnectAttempts; i++ {
				err, ok := q.serveQueue(ctx, client, queue)
				if err != nil {
					log.Printf("%s: error, trying to reconnect: %v", queue, err)

					if ok {
						i = 0
					}
					time.Sleep(3 * time.Second)
					continue
				}
				break
			}
			wg.Done()
		}(queue)
	}
	wg.Wait()

	return nil
}
