package core

import (
	"time"
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
	// chief with more than 1 empleyee
	Business Feature = 1 << iota

	// Creator means actress, writer or artist.
	Creator
)

// PersonConfig represents a minimum piece of information, required for prediction
type PersonConfig struct {
	Name     string    `yaml:"name"        validate:"nonzero"`
	Gender   Gender    `yaml:"gender"      validate:"min=0,max=2"`
	Birthday time.Time `yaml:"birthday"    validate:"nonzero"`
	Features Feature   `yaml:"features"    validate:"min=0,max=3"`
}

// Person contains all the information required for prediction
type Person struct {
	Name          string
	Gender        Gender
	Age           uint8
	Birthday      time.Time
	BaseCards     map[string]*Card
	PersonalCards *PersonalCards

	Rows struct {
		H *Row
		V *Row
	}

	PlanetCycles *PlanetCycles

	Matrix *YearMatrix // Matrix computed based on person age
}

// NewPerson receives valid PersonConfig and valid Locale and
// returns a Person
func NewPerson(conf *PersonConfig, loc *Locale) (*Person, error) {
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
	var pp *Planets
	var plcc *PlanetCycles
	var pcc *PersonalCards
	var mc, dc, sc, pc, rc, lc *Card
	var hr, vr *Row
	var ym *YearMatrix

	// get base primitives from locale
	od = loc.GetOrderedDeck()
	hm = loc.GetHumansMatrix()
	mm = loc.GetYearMatrices()
	cc = loc.GetCycles()
	pp = loc.GetPlanets()

	// get base info from conf
	n = conf.Name
	b = conf.Birthday
	g = conf.Gender
	f = conf.Features
	a = uint8(time.Since(b).Hours() / 24 / 365)

	// In case of Joker, there is nothing to Find. It has only Main card
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

	// Find main cards
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

	if plcc, err = ComputePlanetCycles(b, cc, pp, hr, vr); err != nil {
		return nil, err
	}

	if pcc, err = ComputePersonalCards(mc, g, f, a, loc); err != nil {
		return nil, err
	}

	return &Person{
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

		Rows: struct {
			H *Row
			V *Row
		}{
			hr,
			vr,
		},
		PlanetCycles:  plcc,
		PersonalCards: pcc,
		Matrix:        ym,
	}, nil
}
