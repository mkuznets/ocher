package queue

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

var (
	errNoTasks = errors.New("no tasks found")
	identRe    = regexp.MustCompile(`[^a-zA-Z0-9_]+`)
)

const maxRetries = 3

type Queue struct {
	base *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Queue, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		panic(err)
	}

	return &Queue{
		base: pool,
	}, nil
}

func (q *Queue) Task(ctx context.Context, name string) (*Task, error) {
	for {
		task, err := q.tryGetTask(ctx, name)
		if err == errNoTasks {
			if e := q.waitForTask(ctx, name); e != nil {
				return nil, e
			}
			continue
		} else if err != nil {
			return nil, err
		}
		return task, nil
	}
}

func (q *Queue) tryGetTask(ctx context.Context, name string) (*Task, error) {
	tx, err := q.base.Begin(ctx)

	ok := false
	defer func() {
		if !ok {
			_ = tx.Rollback(ctx)
		}
	}()

	if err != nil {
		return nil, errors.Wrap(err, "could not start transaction")
	}

	var (
		taskID uint64
		args   []byte
	)

	err = tx.QueryRow(ctx, `
	UPDATE "ocher_tasks" SET status='EXECUTING'
	WHERE id=(
		SELECT id FROM "ocher_tasks"
		WHERE status='ENQUEUED' AND name=$1
		ORDER BY id
		FOR UPDATE SKIP LOCKED
		LIMIT 1
	) RETURNING id, args
	;`, name).Scan(&taskID, &args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errNoTasks
		}
		return nil, errors.WithStack(err)
	}

	task := &Task{
		Args: args,
		id:   taskID,
		ctx:  ctx,
		tx:   tx,
		base: q.base,
		name: name,
	}

	var retries int
	err = task.tx.QueryRow(ctx, `
	SELECT COUNT(*) FROM "ocher_statuses" WHERE task_id=$1 AND status='EXECUTING'
	`, taskID).Scan(&retries)

	if retries > maxRetries {
		task.status = "EXPIRED"
		msg := fmt.Sprintf("retry limit reached: %d", maxRetries)
		if err := task.ReportWith(task.tx, "ERROR", msg); err != nil {
			return nil, err
		}

		if err := task.Commit(); err != nil {
			return nil, err
		}
		return nil, errNoTasks
	}
	if err != nil {
		return nil, err
	}

	ok = true

	return task, nil
}

func (q *Queue) waitForTask(ctx context.Context, name string) error {

	waitCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := q.base.Acquire(waitCtx)
	if err != nil {
		return err
	}
	defer conn.Release()

	channelName := identRe.ReplaceAllLiteralString(name, "_")

	if _, err := conn.Exec(waitCtx, fmt.Sprintf("LISTEN ocher_%s;", channelName)); err != nil {
		return err
	}

	_, _ = conn.Conn().WaitForNotification(waitCtx)
	if _, err := conn.Exec(ctx, fmt.Sprintf("UNLISTEN ocher_%s;", channelName)); err != nil {
		return err
	}

	return nil
}
