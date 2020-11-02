package ocher

import "context"

type Task interface {
	ID() uint64
	Name() string
	Args() []byte
	Report(message string) error
}

type TaskFunc func(context.Context, Task) (result []byte, err error)

type taskMeta struct {
	name string
	fn   TaskFunc
}

type taskData struct {
	id     uint64
	name   string
	args   []byte
	stream *taskStreamWorker
}

func (t *taskData) ID() uint64 {
	return t.id
}

func (t *taskData) Name() string {
	return t.name
}

func (t *taskData) Args() []byte {
	return t.args
}

func (t *taskData) Report(message string) error {
	return t.stream.SendReport(message)
}
