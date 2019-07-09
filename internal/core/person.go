package core

import (
	"fmt"
	"log"
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

	PlanetCycles [7]*PlanetCycle
	Environment  []*Person

	Matrices struct {
		Current *YearMatrix
	}
}

// NewPerson returns pointer to a resolved Person, based on PersonConfig
// and Locale
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

	if err = p.resolvePlanetCycles(pc, loc.Planets); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Person) resolveMainCards(od *Deck, hm *Matrix) error {
	var err error
	var idx uint8

	if p.Cards.Main, err = od.FindCardByBirthday(p.Birthday); err != nil {
		return err
	}

	if idx, err = hm.Decks.Main.indexOf(p.Cards.Main.ID); err != nil {
		return err
	}

	if p.Cards.Drain, err = hm.Decks.Drain.FindCardByIndex(idx); err != nil {
		return err
	}

	if p.Cards.Source, err = hm.Decks.Source.FindCardByIndex(idx); err != nil {
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

	if p.Cards.Longterm, err = ym.Matrix.Decks.Main.FindCardByIndex(idx); err != nil {
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

	if p.Cards.Pluto, err = p.Matrices.Current.Decks.Main.FindCardByIndex(idx); err != nil {
		return err
	}

	if idx+1 >= 52 {
		idx = idx - 52
	}

	idx++

	if p.Cards.PlutoResult, err = p.Matrices.Current.Decks.Main.FindCardByIndex(idx); err != nil {
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

func (p *Person) resolvePlanetCycles(pc *[7][54]*PlanetCycle, planets *[7]*Planet) error {
	var r [7]*Card
	var err error

	if r, err = p.Matrices.Current.Decks.Main.CalcHRow(p.Cards.Main); err != nil {
		return err
	}

	if pc == nil || planets == nil {
		return nil
	}

	cpc := GetCurrentPlanetCycles(p.Birthday, pc, planets)

	for i, v := range r {
		if v == nil {
			continue
		}

		p.PlanetCycles[i] = cpc[i]
		p.PlanetCycles[i].Card = v
	}

	return nil
}

// TODO: resolvePersonalCards
// Men:
//   - in spite of age, each man has Jack with the same Suit as his main card,
//     except the case when a man already has Jack with the same Suit
//     as a main card
//   - if a man over 36 years old, he has King with the same Suit
//   - if a man is a business owner with at least 2 employees and more,
//     or a man is a chief (behaves like a chief) he has King with the same Suit
// Women:
//   - women may have up to 3 personal cards at the same time
//   - in spite of age, each woman has Queen with the same Suit as her main card
//     except the case when a woman already has Queen with the same Suit
//     as a main card
//   - if a woman is an actress, a writer or an artist, she has the Jack with
//     the same Suit
//   - if a woman is a business owner with at least 2 employees and more,
//     or a woman is a chief (behaves like a chief) she has King with the same Suit
//   - women younger than 20 years old has Jack with the same Suit
func (p *Person) resolvePersonalCards(pc *PersonConfig, loc *Locale) error {
	return nil
}
