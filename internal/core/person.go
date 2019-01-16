package core

import (
	"fmt"
	"log"
	"time"
)

type Person struct {
	Birthday time.Time
	// Gender
	Environment map[string]*Person
	Matrices    struct {
		Current *YearMatrix
	}

	Cards struct {
		Main        *Card
		Drain       *Card
		Source      *Card
		Longterm    *Card
		Pluto       *Card
		PlutoResult *Card
		Personal    [3]*Card
	}
}

func NewPerson(t time.Time, od *Deck, mm [90]*YearMatrix, om *Matrix, env map[string]time.Time) (*Person, error) {
	var err error

	p := &Person{
		Birthday:    t,
		Environment: make(map[string]*Person, len(env)),
	}

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

	if err = p.resolveEnvironment(t, od, mm, om, env); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Person) resolveMainCards(od *Deck, hm *Matrix) error {
	var err error
	var idx uint8

	idx = uint8(54 - (p.Birthday.Day() + (int(p.Birthday.Month()) * 2)))

	if p.Cards.Main, err = od.GetCardByNumber(idx + 1); err != nil {
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

func (p *Person) resolveLongtermCard(mm [90]*YearMatrix) error {
	var idx uint8
	var err error

	age := time.Now().Year() - p.Birthday.Year()
	ym := mm[age/7]
	delta := uint8((age - age/7))

	fmt.Println(p.Cards.Main)
	if idx, err = ym.Matrix.Decks.Main.indexOf(p.Cards.Main.ID); err != nil {
		return err
	}
	fmt.Println(idx)

	if idx+delta >= 52 {
		idx = idx + delta - 52
	}

	if p.Cards.Longterm, err = ym.Matrix.Decks.Main.GetCardByNumber(idx); err != nil {
		return err
	}

	return nil
}

func (p *Person) resolveCurrentYearMatrix(mm [90]*YearMatrix) error {
	age := time.Now().Year() - p.Birthday.Year()
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

func (p *Person) resolveEnvironment(t time.Time, od *Deck, mm [90]*YearMatrix, om *Matrix, env map[string]time.Time) error {
	var err error
	var person *Person

	if env == nil {
		return nil
	}

	for k, v := range env {
		if person, err = NewPerson(v, od, mm, om, nil); err != nil {
			log.Printf("Unable to build new person when resolving env: %v", err)
		}

		p.Environment[k] = person
	}

	return nil
}
