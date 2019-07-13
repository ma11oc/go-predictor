package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// codeToLevel redirects OK to DEBUG level logging instead of INFO
// This is example how you can log several gRPC code results
func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		// It is DEBUG
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}

// NewUnaryLoggingInterceptor returns unary interceptor
func NewUnaryLoggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	// Shared options for the logger, with a custom gRPC code to log level function.
	o := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}
	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(logger)

	return grpc_zap.UnaryServerInterceptor(logger, o...)
}

// NewStreamLoggingInterceptor returns unary interceptor
func NewStreamLoggingInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	// Shared options for the logger, with a custom gRPC code to log level function.
	o := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}
	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(logger)

	return grpc_zap.StreamServerInterceptor(logger, o...)
}
