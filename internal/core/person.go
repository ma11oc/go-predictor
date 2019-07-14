package core

import (
	"fmt"
	"time"

	"github.com/go-validator/validator"
)

// Gender could be one of Other, Male, Female.type Gender uint8
type Gender uint8

const (
	// Other actually doesn't make any sense for predictions
	Other Gender = iota

	// Male for men
	Male

	// Female for women
	Female
)

// Feature represents qualities of a person which is very important
// for predictions.
// Could be one of:
//   - BusinessOwner 0x01
//   - Creator       0x02
// Several features can be set simultaneously and be checked like:
//   features & (BusinessOwner | Creator)
type Feature uint8

const (
	// Business means businessmen/businesswomen or
	// chief with more than 1 employee
	Business Feature = 1 << iota

	// Creator means actress, writer or artist.
	Creator
)

// PersonalCards is alias to array of 3 Cards
type PersonalCards []*Card

// PersonProfile represents a minimum piece of information, required for prediction
type PersonProfile struct {
	Name     string    `yaml:"name"        validate:"nonzero"`
	Gender   Gender    `yaml:"gender"      validate:"nonzero,min=0,max=2"`
	Birthday time.Time `yaml:"birthday"    validate:"nonzero"`
	Age      int8      `yaml:"age"`
	Features Feature   `yaml:"features"    validate:"min=0,max=3"`
}

// Person contains all the information required for prediction
type Person struct {
	Name     string    `yaml:"name"        validate:"nonzero"`
	Gender   Gender    `yaml:"gender"      validate:"min=1,max=2"`
	Birthday time.Time `yaml:"birthday"    validate:"nonzero"`
	Age      uint8     `yaml:"age"         validate:"min=0"`

	BaseCards map[string]*Card `yaml:"base_cards"       validate:"nonzero,min=6,max=6"`

	PersonalCards *PersonalCards `yaml:"personal_cards" validate:"nonzero,min=0,max=3"`

	Rows *Rows `yaml:"rows"                             validate:"nonzero"`

	PlanetCycles *PlanetCycles `yaml:"planet_cycles"    validate:"nonzero,min=7,max=7"`

	Matrix *YearMatrix // Matrix computed based on person age
}

// NewPerson receives valid PersonProfile and valid Locale
// It returns a Person
func NewPerson(pp *PersonProfile, loc *Locale) (*Person, error) {
	var err error

	var n string
	var b time.Time
	var g Gender
	var f Feature
	var a uint8

	var od *Deck
	var hm *Matrix
	var mm *Matrices
	var cc *Cycles
	var planets *Planets

	var planetCycles *PlanetCycles
	var personalCards *PersonalCards

	var mc, dc, sc, pc, rc, lc *Card
	var hr, vr *Row
	var ym *YearMatrix

	// Validate PersonProfile
	if err = validator.Validate(pp); err != nil {
		return nil, fmt.Errorf("PersonProfile validation error: %v", err)
	}

	// get base primitives from locale
	od = loc.GetOrderedDeck()
	hm = loc.GetHumansMatrix()
	mm = loc.GetYearMatrices()
	cc = loc.GetCycles()

	planets = loc.GetPlanets()

	// get base info from pp
	n = pp.Name
	b = pp.Birthday
	g = pp.Gender
	f = pp.Features

	if pp.Age < 0 {
		a = uint8(time.Since(b).Hours() / 24 / 365)
	} else {
		a = uint8(pp.Age)
	}

	// In case of Joker, there is nothing to compute. It has only Main card
	// with special meaning. So, handle this special case and return struct.
	if b.Month() == time.December && b.Day() == 31 {
		p := &Person{
			Name:     n,
			Gender:   g,
			Age:      a,
			Birthday: b,
		}

		p.BaseCards = map[string]*Card{"main": loc.Exceptions.Joker}

		return p, nil
	}

	if mc, dc, sc, err = ComputeMainCards(b, od, hm); err != nil {
		return nil, err
	}

	if lc, err = ComputeLongtermCard(mm, mc, a); err != nil {
		return nil, err
	}

	if a > 89 {
		ym = mm[a%90]
	} else {
		ym = mm[a]
	}

	if pc, rc, err = ComputePlutoCards(ym, mc); err != nil {
		return nil, err
	}

	if hr, err = ComputeHRow(ym, mc); err != nil {
		return nil, err
	}

	if vr, err = ComputeVRow(ym, mc); err != nil {
		return nil, err
	}

	if planetCycles, err = ComputePlanetCycles(b, cc, planets, hr, vr); err != nil {
		return nil, err
	}

	if personalCards, err = ComputePersonalCards(mc, g, f, a, loc); err != nil {
		return nil, err
	}

	p := &Person{
		Name:     n,
		Gender:   g,
		Age:      a,
		Birthday: b,
		BaseCards: map[string]*Card{
			"main":     mc,
			"drain":    dc,
			"source":   sc,
			"longterm": lc,
			"pluto":    pc,
			"result":   rc,
		},

		Rows: &Rows{
			H: hr,
			V: vr,
		},
		PlanetCycles:  planetCycles,
		PersonalCards: personalCards,
		Matrix:        ym,
	}

	if err = validator.Validate(p); err != nil {
		return nil, fmt.Errorf("Person validation error: %v", err)
	}

	return p, nil
}
