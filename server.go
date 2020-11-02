package ocher

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"mkuznets.com/go/ocher/internal/pb"
	"mkuznets.com/go/ocher/internal/queue"
)

const (
	defaultConnTimeout = 10 * time.Second
)

type Server struct {
	dsn    string
	logger Logger
	grpc   *grpc.Server
}

func NewServer(dsn string, opts ...ServerOption) *Server {
	srv := &Server{
		dsn: dsn,
	}
	for _, opt := range opts {
		opt(srv)
	}

	if srv.grpc == nil {
		srv.grpc = grpc.NewServer(grpc.ConnectionTimeout(defaultConnTimeout))
	}

	if srv.logger == nil {
		srv.logger = &nopLogger{}
	}

	return srv
}

func (srv *Server) Serve(lis net.Listener) error {
	ctx := context.Background()

	q, err := queue.New(ctx, srv.dsn)
	if err != nil {
		return err
	}

	s := &ocherServer{
		queue:  q,
		logger: srv.logger,
	}
	pb.RegisterOcherServer(srv.grpc, s)

	return srv.grpc.Serve(lis)
}
