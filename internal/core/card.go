package core

import (
	"fmt"

	"golang.org/x/text/language"
)

type Meaning struct {
	Keywords    string `yaml:"keywords" validate:"nonzero"`
	Description string `yaml:"description" validate:"nonzero"`
}

// Card represents a simple primitive in matrices
type Card struct {
	ID       uint8  `yaml:"id" validate:"nonzero,min=1,max=52"`
	Rank     string `yaml:"rank" validate:"nonzero,min=1,max=2"`
	Suit     string `yaml:"suit" validate:"nonzero,len=3"`
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
func NewCardFromNumber(n uint8, t language.Tag, ll map[language.Tag]*Locale) (*Card, error) {
	if n <= 0 || n > 52 {
		return nil, fmt.Errorf("Unable to create card: invalid number %v", n)
	}

	return ll[t].GetCardByID(n)
}
