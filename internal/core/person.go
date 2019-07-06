package core

import (
	"fmt"
	"log"
	"time"
)

type Bit uint8

const (
	Male Bit = 1 << iota
	Female
	Other
)

type PersonConfig struct {
	Name        string
	Gender      Bit
	Birthday    time.Time
	Environment []*PersonConfig
}

type Person struct {
	Name     string
	Gender   Bit
	Birthday time.Time `yaml:"birthday" valid:"nonzero"`
	Cards    struct {
		Main        *Card
		Drain       *Card
		Source      *Card
		Longterm    *Card
		Pluto       *Card
		PlutoResult *Card
		Personal    [3]*Card
	}

	PlanetCycles [7]*PlanetCycle
	Environment  []*Person

	Matrices struct {
		Current *YearMatrix
	}
}

func NewPerson(pconf *PersonConfig, loc *Locale) (*Person, error) {
	var err error

	od := loc.Base.od
	om := loc.Base.om
	mm := loc.Base.mm
	pc := loc.Base.pc
	env := pconf.Environment

	p := &Person{
		Name:        pconf.Name,
		Gender:      pconf.Gender,
		Birthday:    pconf.Birthday,
		Environment: make([]*Person, len(env)),
	}

	// In case of Joker, there is nothing to resolve. It has only Main card
	// with special meaning. So, handle this special case and return struct.
	if pconf.Birthday.Month() == time.December && pconf.Birthday.Day() == 31 {
		p.Cards.Main = loc.Exceptions.Joker

		return p, nil
	}

	// resolve person
	if err = p.resolveMainCards(od, om); err != nil {
		return nil, err
	}

	if err = p.resolveLongtermCard(mm); err != nil {
		return nil, err
	}

	if err = p.resolveCurrentYearMatrix(mm); err != nil {
		return nil, err
	}

	if err = p.resolvePlutoCards(); err != nil {
		return nil, err
	}

	if err = p.resolveEnvironment(pconf, loc); err != nil {
		return nil, err
	}

	if err = p.resolvePlanetCycles(pc); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Person) resolveMainCards(od *Deck, hm *Matrix) error {
	var err error
	var idx uint8

	if p.Cards.Main, err = od.GetCardByBirthday(p.Birthday); err != nil {
		return err
	}

	if idx, err = hm.Decks.Main.indexOf(p.Cards.Main.ID); err != nil {
		return err
	}

	if p.Cards.Drain, err = hm.Decks.Drain.GetCardByIndex(idx); err != nil {
		return err
	}

	if p.Cards.Source, err = hm.Decks.Source.GetCardByIndex(idx); err != nil {
		return err
	}

	return nil
}

func (p *Person) resolveLongtermCard(mm *[90]*YearMatrix) error {
	var idx uint8
	var err error

	age := uint8(time.Since(p.Birthday).Hours() / 24 / 365)
	ym := mm[age/7]

	if idx, err = ym.Matrix.Decks.Main.indexOf(p.Cards.Main.ID); err != nil {
		return err
	}

	idx += age%7 + 1
	if idx >= 52 {
		idx -= 52
	}

	if p.Cards.Longterm, err = ym.Matrix.Decks.Main.GetCardByIndex(idx); err != nil {
		return err
	}

	return nil
}

func (p *Person) resolveCurrentYearMatrix(mm *[90]*YearMatrix) error {
	age := uint8(time.Since(p.Birthday).Hours() / 24 / 365)
	ym := mm[age]

	if ym == nil {
		return fmt.Errorf("Unable to resolve current year matrix for age: %v", age)
	}

	p.Matrices.Current = ym

	return nil
}

func (p *Person) resolvePlutoCards() error {
	var idx uint8
	var err error

	if p.Matrices.Current == nil || p.Cards.Main == nil {
		return fmt.Errorf("Unable to resolve pluto card: current year matrix or main card wasn't set")
	}

	if idx, err = p.Matrices.Current.Decks.Main.indexOf(p.Cards.Main.ID); err != nil {
		return err
	}

	if idx+8 >= 52 {
		idx = idx - 52
	}

	idx += 8

	if p.Cards.Pluto, err = p.Matrices.Current.Decks.Main.GetCardByIndex(idx); err != nil {
		return err
	}

	if idx+1 >= 52 {
		idx = idx - 52
	}

	idx++

	if p.Cards.PlutoResult, err = p.Matrices.Current.Decks.Main.GetCardByIndex(idx); err != nil {
		return err
	}

	return nil
}

func (p *Person) resolveEnvironment(pc *PersonConfig, loc *Locale) error {
	var err error
	var person *Person

	env := pc.Environment

	if env == nil {
		return nil
	}

	for i, v := range env {
		if person, err = NewPerson(v, loc); err != nil {
			log.Printf("Unable to build new person when resolving env: %v", err)
		}

		p.Environment[i] = person
	}

	return nil
}

func (p *Person) resolvePlanetCycles(pc *[7][54]*PlanetCycle) error {
	var r [7]*Card
	var err error

	if r, err = p.Matrices.Current.Decks.Main.GetHRow(p.Cards.Main); err != nil {
		return err
	}

	if pc == nil {
		return nil
	}

	cpc := GetCurrentPlanetCycles(p.Birthday, pc)

	for i, v := range r {
		if v == nil {
			continue
		}

		p.PlanetCycles[i] = cpc[i]
		p.PlanetCycles[i].Card = v
	}

	return nil
}
