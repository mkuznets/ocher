package ocher

import "mkuznets.com/go/ocher/internal/pb"

type taskStreamWorker struct {
	stream pb.Ocher_TxTasksClient
}

func (s *taskStreamWorker) Task() (Task, error) {
	in, err := s.stream.Recv()
	if err != nil {
		return nil, err
	}

	return &taskData{
		id:     in.Id,
		name:   in.Name,
		args:   in.Args,
		stream: s,
	}, nil
}

func (s *taskStreamWorker) SendInit(id, task string) error {
	m := &pb.Status{
		Status: &pb.Status_Init{
			Init: &pb.InitStatus{
				Id:       id,
				TaskName: task,
			},
		},
	}
	return s.stream.Send(m)
}

func (s *taskStreamWorker) SendError(err error, data []byte) error {
	m := &pb.Status{
		Status: &pb.Status_Error{
			Error: &pb.ErrorStatus{
				Message: err.Error(),
				Data:    data,
			},
		},
	}
	return s.stream.Send(m)
}

func (s *taskStreamWorker) SendFinish(data []byte) error {
	m := &pb.Status{
		Status: &pb.Status_Finish{
			Finish: &pb.FinishStatus{
				Data: data,
			},
		},
	}
	return s.stream.Send(m)
}

func (s *taskStreamWorker) SendReport(message string) error {
	m := &pb.Status{
		Status: &pb.Status_Report{
			Report: &pb.ReportStatus{
				Message: message,
			},
		},
	}
	return s.stream.Send(m)
}
