package v1

import (
	"context"
	"time"

	// "database/sql"
	// "fmt"
	// "time"

	// "github.com/golang/protobuf/ptypes"
	// "google.golang.org/grpc/codes"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// gw "bitbucket.org/shchukin_a/go-predictor/api"
	core "bitbucket.org/shchukin_a/go-predictor/internal/core"
	// pb "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
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

	locales := core.MustBuildLocales(locPath)

	return &predictorServiceServer{
		Locales: locales,
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

// checkAPI checks if the API version requested by client is supported by server
func (s *predictorServiceServer) getLocale(lang string) (*core.Locale, error) {
	var tag language.Tag
	var err error

	if tag, err = language.Parse(lang); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unknown language '%s'", lang)
	}

	langs := make([]language.Tag, 0, len(s.Locales))

	for k := range s.Locales {
		if tag == k {
			return s.Locales[tag], nil
		}

		langs = append(langs, k)
	}

	return nil, status.Errorf(codes.Unimplemented,
		"unsupported language: service speaks in '%s' languages, but asked for '%s'", langs, tag)
}

func (s *predictorServiceServer) GetGeneralPrediction(ctx context.Context, req *v1.GeneralRequest) (*v1.PredictionResponse, error) {
	var locale *core.Locale
	var person *core.Person
	var err error

	// check if the API version requested by client is supported by server
	if err = s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if locale, err = s.getLocale(req.Lang); err != nil {
		return nil, err
	}

	pc := req.GetPersonConfig()

	b, err := time.Parse("2006-01-02", pc.GetBirthday())
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

	if err = validator.Validate(pconf); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if person, err = core.NewPerson(pconf, locale); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	scs := spew.ConfigState{
		Indent:   "    ",
		MaxDepth: 3,
	}
	scs.Dump(pconf)

	// message PredictionResponse {
	//   string api                        = 1;
	//   string lang                       = 2;
	//
	//   repeated PlanetCycle planetCycles = 3;
	//   map<string, Card> cards           = 4;
	// }

	return &v1.PredictionResponse{
		Api:  apiVersion,
		Lang: req.Lang,

		Cards: map[string]*v1.Card{
			"main": &v1.Card{
				Id:    uint32(person.Cards.Main.ID),
				Rank:  person.Cards.Main.Rank,
				Suit:  person.Cards.Main.Suit,
				Title: person.Cards.Main.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.Cards.Main.Meanings.General.Keywords,
					Description: person.Cards.Main.Meanings.General.Description,
				},
			},
			"drain": &v1.Card{
				Id:    uint32(person.Cards.Drain.ID),
				Rank:  person.Cards.Drain.Rank,
				Suit:  person.Cards.Drain.Suit,
				Title: person.Cards.Drain.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.Cards.Drain.Meanings.General.Keywords,
					Description: person.Cards.Drain.Meanings.General.Description,
				},
			},
			"source": &v1.Card{
				Id:    uint32(person.Cards.Source.ID),
				Rank:  person.Cards.Source.Rank,
				Suit:  person.Cards.Source.Suit,
				Title: person.Cards.Source.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.Cards.Source.Meanings.General.Keywords,
					Description: person.Cards.Source.Meanings.General.Description,
				},
			},
			"longterm": &v1.Card{
				Id:    uint32(person.Cards.Longterm.ID),
				Rank:  person.Cards.Longterm.Rank,
				Suit:  person.Cards.Longterm.Suit,
				Title: person.Cards.Longterm.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.Cards.Longterm.Meanings.Longterm.Keywords,
					Description: person.Cards.Longterm.Meanings.Longterm.Description,
				},
			},
			"pluto": &v1.Card{
				Id:    uint32(person.Cards.Pluto.ID),
				Rank:  person.Cards.Pluto.Rank,
				Suit:  person.Cards.Pluto.Suit,
				Title: person.Cards.Pluto.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.Cards.Longterm.Meanings.Pluto.Keywords,
					Description: person.Cards.Longterm.Meanings.Pluto.Description,
				},
			},
			"pluto/result": &v1.Card{
				Id:    uint32(person.Cards.PlutoResult.ID),
				Rank:  person.Cards.PlutoResult.Rank,
				Suit:  person.Cards.PlutoResult.Suit,
				Title: person.Cards.PlutoResult.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.Cards.Longterm.Meanings.Result.Keywords,
					Description: person.Cards.Longterm.Meanings.Result.Description,
				},
			},
		},
	}, nil
}
