package core

import (
	"fmt"
	"time"
)

// Deck is an unordered deck
type Deck struct {
	Cards [52]*Card `yaml:"cards"`
}

// NewOrderedDeck returns *Deck which contains all the 52 cards,
// arranged by card id
//   See README.md > Appendix > Cards
func NewOrderedDeck(loc *Locale) *Deck {
	var c *Card
	var err error

	od := &Deck{}

	for i := uint8(0); i < 52; i++ {
		if c, err = NewCardFromNumber(i+1, loc); err != nil {
			panic("Unable to create Deck")
		}

		od.Cards[i] = c
	}

	return od
}

// NewDeckFromSlice receives an array with card numbers and returns *Deck
func NewDeckFromSlice(s [52]uint8, od *Deck) (*Deck, error) {
	var err error

	d := &Deck{}

	for i := uint(0); i < 52; i++ {
		if d.Cards[i], err = od.FindCardByNumber(s[i]); err != nil {
			return nil, err
		}
	}
	return d, nil
}

// FindCardByNumber returns appropriate card from a deck
func (d Deck) FindCardByNumber(n uint8) (*Card, error) {
	if n <= 0 || n > 52 {
		return nil, fmt.Errorf("Invalid card number: %d", n)
	}

	for _, c := range d.Cards {
		if c.ID == n {
			return c, nil
		}
	}

	return nil, fmt.Errorf("No such card with number %v was found in the deck", n)
}

// FindCardByBirthday returns appropriate card from a deck
//   See README.md > Appendix > Calendar
func (d Deck) FindCardByBirthday(t time.Time) (*Card, error) {
	idx := 54 - (t.Day() + (int(t.Month()) * 2)) + 1

	if idx <= 0 || idx > 52 {
		return nil, fmt.Errorf("Got invalid card number `%v` for birthday: %v", idx, t)
	}

	return d.FindCardByNumber(uint8(idx))
}

// FindCardByIndex receives index of card in a Deck and returns Card
func (d Deck) FindCardByIndex(i uint8) (*Card, error) {
	var c *Card

	if c = d.Cards[i]; c == nil {
		return nil, fmt.Errorf("No such card with index %v was found in the deck", i)
	}

	return c, nil
}

// indexOf receives number (id) of card and returns its index in a Deck
func (d Deck) indexOf(value uint8) (uint8, error) {
	for i, v := range d.Cards {
		if v.ID == value {
			return uint8(i), nil
		}
	}

	return 255, fmt.Errorf("Unable to find index of `%v`", value)
}

// AsNumbers returns array with cards numbers
func (d Deck) AsNumbers() [52]uint8 {
	var s [52]uint8

	for i, v := range d.Cards {
		s[i] += v.ID
	}

	return s
}

// AsUnicode returns array with cards strings like 'K♥'
func (d Deck) AsUnicode() [52]string {
	var s [52]string

	for i, v := range d.Cards {
		s[i] = fmt.Sprintf("%v%v ", v.Rank, v.Suit)
	}

	return s
}
