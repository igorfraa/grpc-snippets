package main

import (
	"os"
	"strconv"

	"github.com/igorfraa/grpc-snippets/internal/handler"
	"go.uber.org/zap"
)

//go:generate protoc --go_out=../internal/model/grpc --go_opt=paths=source_relative --go-grpc_out=../internal/model/grpc --go-grpc_opt=paths=source_relative --proto_path=../proto example.proto

func main() {
	logger, _ := zap.NewProduction()
	errorChan := make(chan error)
	go func() {
		port, _ := strconv.ParseInt(os.Getenv("SERVICE_PORT"), 10, 64)
		grpcServer := handler.GRPCServer{
			Logger: logger,
			Port:   port,
		}
		logger.Info("starting gRPC at", zap.Any("port", grpcServer.Port))
		errorChan <- grpcServer.Run()
	}()

	select {
	case e := <-errorChan:
		panic(e)
	}
}
