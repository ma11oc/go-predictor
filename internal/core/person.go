package core

import (
	"time"
	// "github.com/davecgh/go-spew/spew"
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

// Features represents qualities of a person which is very important
// for predictions.
// Could be one of:
//   - BusinessOwner 0x01
//   - Creator       0x02
// Several features can be set simultaneously and be checked like:
//   features|BusinessOwner != 0
type Features uint8

const (
	// BusinessOwner means businessmen/businesswomen or
	// chief with more than 1 empleyee
	BusinessOwner Features = 1 << iota

	// Creator means actress, writer or artist.
	Creator
)

// PersonConfig represents a minimum piece of information, required for prediction
type PersonConfig struct {
	Name        string          `yaml:"name"        validate:"nonzero"`
	Gender      Gender          `yaml:"gender"      validate:"min=0,max=2"`
	Birthday    time.Time       `yaml:"birthday"    validate:"nonzero"`
	Features    Features        `yaml:"features"    validate:"min=0,max=3"`
	Environment []*PersonConfig `yaml:"environment"`
}

// Person contains all the information required for prediction
type Person struct {
	Name     string
	Gender   Gender
	Age      uint8
	Birthday time.Time
	Cards    struct {
		Main        *Card
		Drain       *Card
		Source      *Card
		Longterm    *Card
		Pluto       *Card
		PlutoResult *Card
		Personal    [3]*Card
	}

	Rows struct {
		H *Row
		V *Row
	}

	PlanetCycles map[string]*PlanetCycle

	Matrix *YearMatrix // Matrix computed based on person age
}

// NewPerson receives valid PersonConfig and valid Locale and
// returns a Person
func NewPerson(conf *PersonConfig, loc *Locale) (*Person, error) {
	var err error
	var name string
	var birthday time.Time
	var gender Gender
	var age uint8
	var od *Deck
	var om *Matrix
	var mm *Matrices
	// var cc *Cycles
	var mc, dc, sc, pc, rc, lc *Card
	var hr, vr *Row
	var ym *YearMatrix

	// get base primitives from locale
	od = loc.Base.od
	om = loc.Base.om
	mm = loc.Base.mm
	// cc = loc.Base.cc

	// get base info from conf
	name = conf.Name
	birthday = conf.Birthday
	gender = conf.Gender
	age = uint8(time.Since(birthday).Hours() / 24 / 365)

	// In case of Joker, there is nothing to Find. It has only Main card
	// with special meaning. So, handle this special case and return struct.
	if birthday.Month() == time.December && birthday.Day() == 31 {
		p := &Person{
			Name:     name,
			Gender:   gender,
			Age:      age,
			Birthday: birthday,
		}

		p.Cards.Main = loc.Exceptions.Joker

		return p, nil
	}

	// Find main cards
	if mc, dc, sc, err = FindMainCards(birthday, od, om); err != nil {
		return nil, err
	}

	if lc, err = FindLongtermCard(mm, mc, age); err != nil {
		return nil, err
	}

	if age > 89 {
		ym = mm[age%90]
	} else {
		ym = mm[age]
	}

	if pc, rc, err = FindPlutoCards(ym, mc); err != nil {
		return nil, err
	}

	if hr, err = FindHRow(ym, mc); err != nil {
		return nil, err
	}

	if vr, err = FindVRow(ym, mc); err != nil {
		return nil, err
	}

	/*
	 * if err = FindPlanetCycles(pc, loc.Planets); err != nil {
	 *     return nil, err
	 * }
	 */

	return &Person{
		Name:     name,
		Gender:   gender,
		Age:      age,
		Birthday: birthday,
		Cards: struct {
			Main        *Card
			Drain       *Card
			Source      *Card
			Longterm    *Card
			Pluto       *Card
			PlutoResult *Card
			Personal    [3]*Card
		}{
			mc,
			dc,
			sc,
			lc,
			pc,
			rc,
			[3]*Card{},
		},

		Rows: struct {
			H *Row
			V *Row
		}{
			hr,
			vr,
		},
	}, nil
}
