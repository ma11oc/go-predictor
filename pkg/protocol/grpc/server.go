package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	v1 "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
	"bitbucket.org/shchukin_a/go-predictor/pkg/logger"
	"bitbucket.org/shchukin_a/go-predictor/pkg/protocol/grpc/middleware"
)

// RunServer runs gRPC service to publish Predictor service
func RunServer(ctx context.Context, v1API v1.PredictorServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterPredictorServiceServer(server, v1API)

	// Register reflection service on gRPC server.
	// FIXME: must be removed on prod
	reflection.Register(server)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
