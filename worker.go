package ocher

import (
	"context"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
	"mkuznets.com/go/ocher/internal/pb"
)

const (
	dialTimeout       = 30
	reconnectAttempts = 3
	reconnectDelay    = 3 * time.Second
)

type Worker struct {
	addr   string
	id     string
	tasks  []*taskMeta
	logger Logger
}

func NewWorker(addr string, opts ...WorkerOption) *Worker {
	wrk := &Worker{
		addr:  addr,
		tasks: make([]*taskMeta, 0),
	}
	for _, opt := range opts {
		opt(wrk)
	}

	if wrk.logger == nil {
		wrk.logger = &nopLogger{}
	}

	return wrk
}

func (wrk *Worker) RegisterTask(name string, fn TaskFunc) {
	wrk.tasks = append(wrk.tasks, &taskMeta{
		name: name,
		fn:   fn,
	})
}

func (wrk *Worker) Serve() (err error) {
	ctx := context.Background()

	dialCtx, cancel := context.WithTimeout(ctx, dialTimeout*time.Second)
	defer cancel()

	wrk.logger.Log(ctx, LogLevelDebug, "dialing gRPC target", map[string]interface{}{
		"addr": wrk.addr,
	})

	conn, err := grpc.DialContext(dialCtx, wrk.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	wrk.logger.Log(ctx, LogLevelDebug, "connected to gRPC target", map[string]interface{}{
		"addr": wrk.addr,
	})

	defer func() {
		if e := conn.Close(); e != nil {
			err = e
			return
		}
	}()

	client := pb.NewOcherClient(conn)

	var wg sync.WaitGroup

	for _, meta := range wrk.tasks {
		wg.Add(1)

		go func(tm *taskMeta) {
			for i := 0; i < reconnectAttempts; i++ {
				err, ok := wrk.serveTask(ctx, client, tm)
				if err != nil {
					wrk.logger.Log(ctx, LogLevelError, "error, trying to reconnect", map[string]interface{}{
						"task": tm.name,
						"err":  err,
					})

					if ok {
						i = 0
					}
					time.Sleep(reconnectDelay)
					continue
				}
				break
			}
			wg.Done()
		}(meta)
	}
	wg.Wait()

	return nil
}

func (wrk *Worker) serveTask(ctx context.Context, client pb.OcherClient, tm *taskMeta) (error, bool) {

	streamCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := client.TxTasks(streamCtx, grpc.WaitForReady(true))
	if err != nil {
		return err, false
	}
	stream := &taskStreamWorker{s}

	wrk.logger.Log(ctx, LogLevelDebug, "connected to queue", map[string]interface{}{
		"task": tm.name,
	})

	if err := stream.SendInit(wrk.id, tm.name); err != nil {
		return err, false
	}

	wrk.logger.Log(ctx, LogLevelInfo, "queue initialised, waiting for tasks", map[string]interface{}{
		"task": tm.name,
	})

	for {
		task, err := stream.Task()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err, true
		}

		wrk.logger.Log(ctx, LogLevelInfo, "new task", map[string]interface{}{
			"id": task.ID(),
		})

		result, taskErr := tm.fn(ctx, task)

		if taskErr != nil {
			wrk.logger.Log(ctx, LogLevelError, "task error", map[string]interface{}{
				"id":  task.ID(),
				"err": taskErr,
			})
			if err := stream.SendError(taskErr, result); err != nil {
				return err, true
			}
		} else {
			wrk.logger.Log(ctx, LogLevelInfo, "task finished", map[string]interface{}{
				"id": task.ID(),
			})
			if err := stream.SendFinish(result); err != nil {
				return err, true
			}
		}
	}
}
