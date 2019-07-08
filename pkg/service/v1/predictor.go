package v1

import (
	"context"
	"time"

	// "database/sql"
	// "fmt"
	// "time"

	// "github.com/golang/protobuf/ptypes"
	// "google.golang.org/grpc/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// gw "bitbucket.org/shchukin_a/go-predictor/api"
	core "bitbucket.org/shchukin_a/go-predictor/internal/core"
	pb "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-validator/validator"

	// "bitbucket.org/shchukin_a/go-predictor/pkg/logger"

	v1 "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
	// "github.com/golang/glog"
	// "github.com/golang/protobuf/ptypes/empty"
	// "github.com/golang/protobuf/ptypes/struct"
	// "go.uber.org/zap"
	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type predictorServiceServer struct {
	Locales core.Locales
}

// NewPredictorServiceServer creates Predictor service
func NewPredictorServiceServer(locPath string) v1.PredictorServiceServer {
	return &predictorServiceServer{
		Locales: core.BuildLocales(locPath),
	}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *predictorServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *predictorServiceServer) FindCardByBirthday(ctx context.Context, req *v1.CardRequest) (*pb.CardResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	pc := req.GetPersonConfig()

	b, err := time.Parse("2006-01-01", pc.GetBirthday())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	pconf := &core.PersonConfig{
		Name:     pc.GetName(),
		Gender:   core.Gender(pc.GetGender()),
		Birthday: b,
		Features: core.Features(pc.GetFeatures()),
		// Environment: in.GetEnvironment(),
	}

	if err := validator.Validate(pconf); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	scs := spew.ConfigState{
		Indent:   "    ",
		MaxDepth: 3,
	}
	scs.Dump(pconf)

	return &v1.CardResponse{}, nil
}
