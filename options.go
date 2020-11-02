package ocher

type ServerOption = func(*Server)

func ServerLogger(logger Logger) ServerOption {
	return func(w *Server) {
		w.logger = logger
	}
}

type WorkerOption = func(*Worker)

func WorkerID(id string) WorkerOption {
	return func(w *Worker) {
		w.id = id
	}
}

func WorkerLogger(logger Logger) WorkerOption {
	return func(w *Worker) {
		w.logger = logger
	}
}
