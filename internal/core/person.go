package core

import (
	"fmt"
	"time"
)

type Person struct {
	Birthday    time.Time
	Environment []*Person
	Matrices    struct {
		Current *YearMatrix
	}

	Cards struct {
		Main     *Card
		Drain    *Card
		Source   *Card
		Longterm *Card
		Pluto    *Card
		Result   *Card
		Personal [3]*Card
	}
}

func NewPerson(t time.Time, od *Deck, mm [90]*YearMatrix, om *Matrix) (*Person, error) {
	var err error

	p := &Person{
		Birthday: t,
	}

	if err = p.resolveMainCards(od, om); err != nil {
		return nil, err
	}

	if err = p.resolveLongtermCard(mm); err != nil {
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

	fmt.Println(p.Cards.Main)
	if idx, err = ym.Matrix.Decks.Main.indexOf(p.Cards.Main.ID); err != nil {
		return err
	}
	fmt.Println(idx)

	if idx >= 52 {
		idx -= 52
	}

	if p.Cards.Longterm, err = ym.Matrix.Decks.Main.GetCardByNumber(idx); err != nil {
		return err
	}

	return nil
}
