package core

import (
	"fmt"
)

type Meaning struct {
	Keywords    string `yaml:"keywords" validate:"nonzero"`
	Description string `yaml:"description" validate:"nonzero"`
}

// Card represents a simple primitive in matrices
type Card struct {
	ID       uint8  `yaml:"id" validate:"min=0,max=52"`
	Rank     string `yaml:"rank" validate:"nonzero,regexp=^(A|2|3|4|5|6|7|8|9|10|J|Q|K|Joker)$"`
	Suit     string `yaml:"suit" validate:"nonzero,min=3,max=4"`
	Title    string `yaml:"title" validate:"nonzero"`
	Meanings struct {
		General  Meaning
		Longterm Meaning
		Mercury  Meaning
		Venus    Meaning
		Mars     Meaning
		Jupiter  Meaning
		Saturn   Meaning
		Uranus   Meaning
		Neptune  Meaning
		Pluto    Meaning
		Result   Meaning
	}
}

// NewCardFromNumber returns type *Card from the given number (0, 52]
func NewCardFromNumber(n uint8, loc *Locale) (*Card, error) {
	if n <= 0 || n > 52 {
		return nil, fmt.Errorf("Unable to create card: invalid number %v", n)
	}

	return loc.GetCardByID(n)
}
