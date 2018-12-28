package core

import "fmt"

// Deck is an unordered deck
type Deck struct {
	Cards [52]*Card
}

/*
 * func (d Deck) String() string {
 *     var s string
 *
 *     for i, v := range d.Cards {
 *         s += fmt.Sprintf("%v ", v.Number)
 *
 *         if i != len(d.Cards)-1 {
 *             s += " "
 *         }
 *     }
 *
 *     return fmt.Sprintf("Deck{cards: [%v]}", s)
 * }
 *
 */
// NewOrderedDeck returns *Deck which contains all the 52 cards,
// in the following order:
//   index:     0  1      12 13     25     38     51
//   card:      A♥ 2♥ ... K♥ A♣ ... A♦ ... A♠ ... K♠
func NewOrderedDeck() *Deck {
	var c *Card
	var err error

	od := &Deck{}

	for i := uint8(0); i < 52; i++ {
		if c, err = NewCardFromNumber(i + 1); err != nil {
			panic("Unable to create Deck")
		}

		od.Cards[i] = c
	}

	return od
}

func NewDeckFromSlice(s [52]uint8) (*Deck, error) {
	var err error

	d := &Deck{}

	for i := uint(0); i < 52; i++ {
		if d.Cards[i], err = NewCardFromNumber(s[i]); err != nil {
			return nil, err
		}
	}
	return d, nil
}

// GetCardByNumber returns appropriate card from a deck
func (d Deck) GetCardByNumber(n uint8) (*Card, error) {
	if n <= 0 || n > 52 {
		return nil, fmt.Errorf("Invalid card number: %d", n)
	}

	for _, c := range d.Cards {
		if c.Number == n {
			return c, nil
		}
	}

	return nil, fmt.Errorf("No such card with number %v was found in the deck", n)
}

func (d *Deck) indexOf(value uint8) (uint8, error) {
	for i, v := range d.Cards {
		if v.Number == value {
			return uint8(i), nil
		}
	}

	return 255, fmt.Errorf("Unable to find index of `%v`", value)
}
