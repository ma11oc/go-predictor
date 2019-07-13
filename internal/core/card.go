package core

import (
	"time"
)

var (
	// BaseCardsOrder sets the order of base cards. Useful in iterations over
	// a Person cards
	BaseCardsOrder = [6]string{"main", "drain", "source", "longterm", "pluto", "result"}
)

// Meaning describes a card
type Meaning struct {
	Keywords    string `yaml:"keywords"    validate:"nonzero"`
	Description string `yaml:"description" validate:"nonzero"`
}

// Card represents a simple primitive in matrices
type Card struct {
	ID       uint8  `yaml:"id"    validate:"min=0,max=52"`
	Rank     string `yaml:"rank"  validate:"nonzero,regexp=^(A|2|3|4|5|6|7|8|9|10|J|Q|K|Joker)$"`
	Suit     string `yaml:"suit"  validate:"nonzero,min=3,max=4,regexp=^(♥|♣|♦|♠|🃏)$"`
	Title    string `yaml:"title" validate:"nonzero"`
	Meta     string
	Meanings map[string]Meaning `yaml:"meanings" validate:"nonzero,min=11,max=11"`
}

// PersonalCards is alias to array of 3 Cards
type PersonalCards [3]*Card

// GetBirthdays returns array with all the birthday dates during a year,
// associated with particular card
func (c Card) GetBirthdays() ([]time.Time, error) {
	var bdays []time.Time

	for m := 1; m <= 12; m++ {
		day := 54 - int(c.ID) - 2*(m-1) - 1

		days := time.Date(2000, time.Month(m), 1, 0, 0, 0, 0, time.UTC).
			AddDate(0, 1, 0).
			Sub(time.Date(2000, time.Month(m), 1, 0, 0, 0, 0, time.UTC))

		if day >= 0 && day <= int(days.Hours()/24) {
			bdays = append(bdays, time.Date(2000, time.Month(m), day, 0, 0, 0, 0, time.UTC))
		}
	}

	return bdays, nil
}

// NewCardFromNumber returns type *Card from the given number (0, 52]
func NewCardFromNumber(n uint8, loc *Locale) (*Card, error) {
	return loc.FindCardByID(n)
}

// NewCardFromString returns card, found by string like 'K♠'
func NewCardFromString(s string, loc *Locale) (*Card, error) {
	return loc.FindCardByString(s)
}
