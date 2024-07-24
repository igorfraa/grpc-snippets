package main

import (
	"log"

	"github.com/igorfraa/grpc-snippets/internal/config"
	"github.com/igorfraa/grpc-snippets/internal/handler"
	"go.uber.org/zap"
)

//nolint:lll
//go:generate protoc --go_out=../../internal/model/grpc --go_opt=paths=source_relative --go-grpc_out=../../internal/model/grpc --go-grpc_opt=paths=source_relative --proto_path=../../proto ping.proto
//go:generate protoc --plugin=protoc-gen-grpc-web=../../frontend/ping-example/node_modules/.bin/protoc-gen-grpc-web --grpc-web_out=import_style=commonjs+dts,mode=grpcweb:../../frontend/ping-example/src/lib/ --js_out=import_style=commonjs:../../frontend/ping-example/src/lib/ -I ../../proto ping.proto
func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logger, _ := zap.NewProduction()
	errorChan := make(chan error)
	go func() {
		grpcServer := handler.GRPCServer{
			Logger: logger,
			Port:   c.GRPCPort,
		}
		logger.Info("starting gRPC at", zap.Any("port", grpcServer.Port))
		errorChan <- grpcServer.Run()
	}()

	select {
	case e := <-errorChan:
		panic(e)
	}
}
