package handler

import (
	"context"
	"fmt"
	"net"

	pb "github.com/igorfraa/grpc-snippets/internal/model/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedRootServer
	Logger *zap.Logger
	Port   int64
}

func (s *GRPCServer) Run() error {
	var (
		err    error
		listen net.Listener
		opts   []grpc.ServerOption
	)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRootServer(grpcServer, s)
	if listen, err = net.Listen("tcp", fmt.Sprintf(":%d", s.Port)); err != nil {
		s.Logger.Fatal("Can not attach on network address",
			zap.Any("port", s.Port),
			zap.Error(err),
		)
	}
	return grpcServer.Serve(listen)
}

func (s *GRPCServer) Ping(_ context.Context, req *pb.PingMessage) (*pb.PingMessage, error) {
	s.Logger.Debug("got gRPC ping", zap.Int64("seq", req.GetSeq()), zap.Any("message", req.GetPayload()))
	pingResponse := "pong"
	req.Payload = &pingResponse
	return req, nil
}
