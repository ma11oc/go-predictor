package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

// NewUnaryValidatorInterceptor returns unary validator
func NewUnaryValidatorInterceptor() grpc.UnaryServerInterceptor {
	return grpc_validator.UnaryServerInterceptor()
}

// NewStreamValidatorInterceptor returns stream validator
func NewStreamValidatorInterceptor() grpc.StreamServerInterceptor {
	return grpc_validator.StreamServerInterceptor()
}
