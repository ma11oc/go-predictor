package core

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	ranks = [13]string{
		"A",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"J",
		"Q",
		"K",
	}

	// TODO: suits as runes
	suits = [4]string{
		"\u2665", // ♥ hearts
		"\u2663", // ♣ clovers
		"\u2666", // ♦ tiles
		"\u2660", // ♠ pikes
	}
)

type Meaning struct {
	Keywords    string `yaml:"keywords"`
	Description string `yaml:"description"`
}

// Card represents a simple priimtive in matrices
type Card struct {
	ID       uint8  `yaml:"id"`
	Rank     string `yaml:"rank"`
	Suit     string `yaml:"suit"`
	Title    string `yaml:"title"`
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

func (c Card) Localize(lang string) (*Card, error) {
	// FIXME: language matcher
	lt := language.Make(lang)
	p := message.NewPrinter(lt)

	card := &Card{
		ID:    c.ID,
		Suit:  c.Suit,
		Rank:  c.Rank,
		Title: p.Sprintf(strconv.Itoa(int(c.ID)) + ".title"),
		Meanings: struct {
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
		}{
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.general.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.general.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.longterm.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.longterm.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.mercury.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.mercury.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.venus.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.venus.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.mars.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.mars.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.jupiter.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.jupiter.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.saturn.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.saturn.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.uranus.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.uranus.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.neptune.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.neptune.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.pluto.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.pluto.description"),
			},
			Meaning{
				Keywords:    p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.result.keywords"),
				Description: p.Sprintf(strconv.Itoa(int(c.ID)) + ".meanings.result.description"),
			},
		},
	}

	return card, nil
}

// NewCardFromNumber returns type *Card from the given number (0, 52]
func NewCardFromNumber(n uint8, lang string) (*Card, error) {
	if n <= 0 || n > 52 {
		return nil, fmt.Errorf("Unable to create card: invalid number -> %v", n)
	}

	c := &Card{
		ID:   uint8(n),
		Suit: suits[(n-1)/13],
		Rank: ranks[(n-1)%13],
	}

	return c.Localize(lang)
}
