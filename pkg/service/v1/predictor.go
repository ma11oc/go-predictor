package v1

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "bitbucket.org/shchukin_a/go-predictor/internal/core"

	"github.com/go-validator/validator"

	v1 "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
	"bitbucket.org/shchukin_a/go-predictor/pkg/logger"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type predictorServiceServer struct {
	Locales core.Locales
}

// HandlePanic logs a error via zap.Logger
func HandlePanic(f string, logger *zap.Logger) {
	var err string

	if r := recover(); r != nil {
		// find out exactly what the error was and set err
		switch x := r.(type) {
		case string:
			break
		case error:
			err = x.Error()
		default:
			// Fallback err (per specs, error strings should be lowercase w/o punctuation
			err = "unknown panic"
		}

		logger.Error("error", zap.String("msg", fmt.Sprintf("recovered in %v: %v", f, err)))
	}
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

	defer HandlePanic("ComputePerson", logger.Log)

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
		Features: core.Feature(pc.GetFeatures()),
	}

	if err = validator.Validate(pconf); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if person, err = core.NewPerson(pconf, locale); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	// set cards
	baseCards := map[string]*v1.Card{}
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

		baseCards[v] = &v1.Card{
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
	planetCycles := map[string]*v1.PlanetCycle{}

	for i, v := range core.PlanetsOrder {
		planetCycles[v] = &v1.PlanetCycle{
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

	personalCards := []*v1.Card{}

	for _, v := range person.PersonalCards {
		if v != nil {
			personalCards = append(personalCards, &v1.Card{
				Id:    uint32(v.ID),
				Rank:  v.Rank,
				Suit:  v.Suit,
				Title: v.Title,
				Meaning: &v1.Meaning{
					Keywords:    v.Meanings["general"].Keywords,
					Description: v.Meanings["general"].Description,
				},
			})
		}
	}

	return &v1.PersonResponse{
		Api:  apiVersion,
		Lang: req.Lang,

		Person: &v1.Person{
			BaseCards:     baseCards,
			PlanetCycles:  planetCycles,
			PersonalCards: personalCards,
		},
	}, nil
}
