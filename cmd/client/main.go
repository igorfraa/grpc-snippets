package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/igorfraa/grpc-snippets/internal/config"
	pb "github.com/igorfraa/grpc-snippets/internal/model/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	_ = godotenv.Load()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	var (
		ctx        = context.Background()
		remote     = fmt.Sprintf("localhost:%d", c.GRPCPort)
		conn       *grpc.ClientConn
		clientCert credentials.TransportCredentials
		opts       []grpc.DialOption
		callOpts   []grpc.CallOption
		res        *pb.PingMessage
	)

	logger.Debug(remote)
	if c.GRPCClientUseTLS {
		if clientCert, err = getTransportCredentials(c); err != nil {
			log.Fatal(err)
		}
	} else {
		clientCert = insecure.NewCredentials()
		log.Println("using insecure credentials")
	}

	opts = append(opts, grpc.WithTransportCredentials(clientCert))
	log.Println("gRPC client connecting to remote ", remote, clientCert.Info())

	if conn, err = grpc.Dial(remote, opts...); err != nil {
		log.Fatal(err)
	}
	pbClient := pb.NewRootClient(conn)
	req := pb.PingMessage{Payload: "ping", Seq: 0}
	ticker := time.NewTicker(1 * time.Second)
	var t time.Time
	for ; true; t = <-ticker.C {
		res, err = pbClient.Ping(ctx, &req, callOpts...)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%s] Upstream ping response [%d]: [%s]\n", t, res.GetSeq(), res.GetPayload())

		req.Seq++
	}
}

func getTransportCredentials(c *config.Config) (credentials.TransportCredentials, error) {
	var (
		clientCert credentials.TransportCredentials
		err        error
	)

	if c.GRPCClientCertFN == "" {
		log.Println("Using TLS client without verification")
		//nolint:gosec
		clientCert = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	} else {
		log.Println("Using TLS client")
		clientCert, err = credentials.NewClientTLSFromFile("/opt/homebrew/etc/nginx/ssl/fullchain.pem", "")
		if err != nil {
			log.Fatal(err)
		}
	}
	return clientCert, err
}
