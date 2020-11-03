package ocher

import (
	"context"
	"io"

	"github.com/pkg/errors"
	"mkuznets.com/go/ocher/internal/pb"
	"mkuznets.com/go/ocher/internal/queue"
)

type ocherServer struct {
	pb.UnimplementedOcherServer
	logger Logger
	queue  *queue.Queue
}

func (server *ocherServer) TxTasks(stream pb.Ocher_TxTasksServer) error {
	ts := &taskStreamServer{stream: stream}

	ctx := stream.Context()

	init, err := ts.WaitInit()
	if err != nil {
		server.logger.Log(ctx, LogLevelError, "init error", map[string]interface{}{
			"err": err,
		})
		return err
	}

	err = func() error {
		server.logger.Log(ctx, LogLevelInfo, "worker connected", map[string]interface{}{
			"task": init.TaskName,
		})

		for {
			task, err := server.queue.Task(ctx, init.TaskName)
			if err != nil {
				return err
			}

			err = func() error {
				defer task.Close()

				if err := task.InsertStatus("EXECUTING"); err != nil {
					return err
				}

				if err := ts.Execute(task); err != nil {
					return err
				}

				if err := task.Commit(); err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}
		}
	}()

	if err != nil {
		if ctx.Err() == context.Canceled {
			server.logger.Log(ctx, LogLevelInfo, "worker disconnected", map[string]interface{}{
				"task": init.TaskName,
			})
		} else {
			server.logger.Log(ctx, LogLevelError, "queue error", map[string]interface{}{
				"err":  err,
				"task": init.TaskName,
			})
		}
	}

	return err
}

type initData struct {
	ID       string
	TaskName string
}

type taskStreamServer struct {
	stream pb.Ocher_TxTasksServer
}

func (s *taskStreamServer) Execute(task *queue.Task) error {
	taskProto := &pb.Task{
		Id:   task.ID(),
		Name: task.Name(),
		Args: task.Args,
	}
	if err := s.stream.Send(taskProto); err != nil {
		return err
	}

	for {
		// Get task status
		msg, err := s.stream.Recv()

		switch {
		case err == io.EOF:
			return err
		case err != nil:
			return errors.WithStack(err)
		}

		switch x := msg.Status.(type) {

		case *pb.Status_Report:
			if err := task.Report(x.Report.Message); err != nil {
				return err
			}

		case *pb.Status_Error:
			task.Result = cloneBytes(x.Error.Data)
			if err := task.ReportWith(task.Tx(), "ERROR", x.Error.Message); err != nil {
				return err
			}
			task.Fail()
			return nil

		case *pb.Status_Finish:
			task.Result = cloneBytes(x.Finish.Data)
			task.Succeed()
			return nil

		default:
			return errors.New("invalid status message")
		}
	}
}

func (s *taskStreamServer) WaitInit() (*initData, error) {
	msg, err := s.stream.Recv()
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return nil, err
	}

	initMsg := msg.GetInit()
	if initMsg == nil {
		return nil, errors.New("unexpected initial message")
	}

	if initMsg.Id == "" {
		return nil, errors.New("init.id is not defined")
	}
	if initMsg.TaskName == "" {
		return nil, errors.New("init.task_name is not defined")
	}

	init := &initData{
		ID:       initMsg.Id,
		TaskName: initMsg.TaskName,
	}

	return init, nil
}
