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

func (s *predictorServiceServer) ComputePerson(ctx context.Context, req *v1.PersonRequest) (*v1.PersonResponse, error) {
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
	}

	if err = validator.Validate(pconf); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if person, err = core.NewPerson(pconf, locale); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	// set cards
	cc := map[string]*v1.Card{}
	descriptors := locale.Descriptors

	for _, v := range core.BaseCardsOrder {
		kws := ""
		dsc := ""

		if v == "main" || v == "drain" || v == "source" {
			kws = person.BaseCards["main"].Meanings["general"].Keywords
			dsc = person.BaseCards["main"].Meanings["general"].Description
		} else {
			kws = person.BaseCards[v].Meanings[v].Keywords
			dsc = person.BaseCards[v].Meanings[v].Description
		}

		cc[v] = &v1.Card{
			Id:    uint32(person.BaseCards[v].ID),
			Rank:  person.BaseCards[v].Rank,
			Suit:  person.BaseCards[v].Suit,
			Title: person.BaseCards[v].Title,
			Meta:  descriptors.Cards[v],
			Meaning: &v1.Meaning{
				Keywords:    kws,
				Description: dsc,
			},
		}
	}

	// set planet cycles
	pcc := map[string]*v1.PlanetCycle{}

	for i, v := range core.PlanetsOrder {
		pcc[v] = &v1.PlanetCycle{
			Card: &v1.Card{
				Id:    uint32(person.PlanetCycles[i].Cards.H.ID),
				Rank:  person.PlanetCycles[i].Cards.H.Rank,
				Suit:  person.PlanetCycles[i].Cards.H.Suit,
				Title: person.PlanetCycles[i].Cards.H.Title,
				Meaning: &v1.Meaning{
					Keywords:    person.PlanetCycles[i].Cards.H.Meanings[core.PlanetsOrder[i]].Keywords,
					Description: person.PlanetCycles[i].Cards.H.Meanings[core.PlanetsOrder[i]].Description,
				},
			},
			Planet: &v1.Planet{
				Id:     uint32(person.PlanetCycles[i].Planet.ID),
				Name:   person.PlanetCycles[i].Planet.Name,
				Symbol: person.PlanetCycles[i].Planet.Symbol,
			},
			Start: &v1.PlanetCycleDate{
				Month: uint32(person.PlanetCycles[i].Start.Month()),
				Day:   uint32(person.PlanetCycles[i].Start.Day()),
			},
			End: &v1.PlanetCycleDate{
				Month: uint32(person.PlanetCycles[i].End.Month()),
				Day:   uint32(person.PlanetCycles[i].End.Day()),
			},
		}
	}

	return &v1.PersonResponse{
		Api:  apiVersion,
		Lang: req.Lang,

		Person: &v1.Person{
			BaseCards:    cc,
			PlanetCycles: pcc,
		},
	}, nil
}
