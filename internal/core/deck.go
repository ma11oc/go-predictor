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
// in the following order (index: number rank suit):
//    0:  1  A♥ |   1:  2  2♥ |   2:  3  3♥ |   3:  4  4♥ |   4:  5  5♥ |   5:  6  6♥ |   6:  7  7♥ |   7:  8  8♥ |   8:  9  9♥ |   9: 10  10♥ |  10: 11  J♥ |  11: 12  Q♥ |  12: 13  K♥ |
//   13: 14  A♣ |  14: 15  2♣ |  15: 16  3♣ |  16: 17  4♣ |  17: 18  5♣ |  18: 19  6♣ |  19: 20  7♣ |  20: 21  8♣ |  21: 22  9♣ |  22: 23  10♣ |  23: 24  J♣ |  24: 25  Q♣ |  25: 26  K♣ |
//   26: 27  A♦ |  27: 28  2♦ |  28: 29  3♦ |  29: 30  4♦ |  30: 31  5♦ |  31: 32  6♦ |  32: 33  7♦ |  33: 34  8♦ |  34: 35  9♦ |  35: 36  10♦ |  36: 37  J♦ |  37: 38  Q♦ |  38: 39  K♦ |
//   39: 40  A♠ |  40: 41  2♠ |  41: 42  3♠ |  42: 43  4♠ |  43: 44  5♠ |  44: 45  6♠ |  45: 46  7♠ |  46: 47  8♠ |  47: 48  9♠ |  48: 49  10♠ |  49: 50  J♠ |  50: 51  Q♠ |  51: 52  K♠ |
// FIXME: get default locale from config
func NewOrderedDeck(lang string) *Deck {
	var c *Card
	var err error

	od := &Deck{}

	for i := uint8(0); i < 52; i++ {
		if c, err = NewCardFromNumber(i+1, lang); err != nil {
			panic("Unable to create Deck")
		}

		od.Cards[i] = c
	}

	return od
}

func NewDeckFromSlice(s [52]uint8, od *Deck) (*Deck, error) {
	var err error

	d := &Deck{}

	for i := uint(0); i < 52; i++ {
		if d.Cards[i], err = od.GetCardByNumber(s[i]); err != nil {
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
		if c.ID == n {
			return c, nil
		}
	}

	return nil, fmt.Errorf("No such card with number %v was found in the deck", n)
}

// GetCardByBirthday returns appropriate card from a deck
func (d Deck) GetCardByBirthday(t time.Time) (*Card, error) {
	idx := uint8(54-(t.Day()+(int(t.Month())*2))) + 1

	if idx < 1 || idx > 52 {
		return nil, fmt.Errorf("Got invalid card number `%v` for birthday: %v", idx, t)
	}

	return d.GetCardByNumber(idx)
}

// GetCardByIndex returns appropriate card from a deck
func (d Deck) GetCardByIndex(i uint8) (*Card, error) {
	var c *Card

	if c = d.Cards[i]; c == nil {
		return nil, fmt.Errorf("No such card with index %v was found in the deck", i)
	}

	return c, nil
}

func (d Deck) indexOf(value uint8) (uint8, error) {
	for i, v := range d.Cards {
		if v.ID == value {
			return uint8(i), nil
		}
	}

	return 255, fmt.Errorf("Unable to find index of `%v`", value)
}

func (d Deck) AsNumbers() [52]uint8 {
	var s [52]uint8

	for i, v := range d.Cards {
		s[i] += v.ID
	}

	return s
}

func (d Deck) AsUnicode() [52]string {
	var s [52]string

	for i, v := range d.Cards {
		s[i] = fmt.Sprintf("%v%v ", v.Rank, v.Suit)
	}

	return s
}

func (d Deck) GetHRow(c *Card) ([7]*Card, error) {
	var row [7]*Card
	var ci uint8 // card index
	var err error

	if ci, err = d.indexOf(c.ID); err != nil {
		return row, err
	}

	for i := uint8(1); i <= 7; i++ {
		cur := ci + i
		if cur >= 52 {
			cur -= 52
		}

		row[i-1] = d.Cards[cur]
	}

	return row, nil
}

// GetVRow ...
// 6: [48, 41, 34, 27, 20, 13, 6]      [-1, -8,  -15, ...]
// 5: [47, 40, 33, 26, 19, 12, 5]      [-2, -9,  -16, ...]
// 4: [46, 39, 32, 25, 18, 11, 4, 51]  [-3, -10, -17, ...]
// 3: [45, 38, 31, 24, 17, 10, 3, 50]  [-4, -11, -18, ...]
// 2: [44, 37, 30, 23, 16, 9,  2, 49]  [-5, -12, -19, ...]
// 1: [43, 36, 29, 22, 15, 8,  1]      [-6, -13, -20, ...]
// 0: [42, 35, 28, 21, 14, 7,  0]      [-7, -14, -21, ...]
func (d Deck) GetVRow(c *Card) ([7]*Card, error) {
	var row [7]*Card
	var ci uint8 // card index
	var err error

	if ci, err = d.indexOf(c.ID); err != nil {
		return row, err
	}

	for i := 1; i <= 7; i++ {
		cur := int(ci) - i*7
		if cur < 0 {
			switch {
			case cur%7 == -3 || cur%7 == -4 || cur%7 == -5: // we are in
				if cur == -3 || cur == -4 || cur == -5 {
					cur = 54 + cur
				} else {
					cur = 56 + cur
				}
			default:
				if i == 7 {
					continue
				}
				cur = 49 + cur
			}
		}
		// fmt.Printf("ci: %2v, i: %v , cur: %3v | ", ci, i, cur)
		// fmt.Printf("%v: %2v | %v%v\n", i, d.Cards[cur].ID, d.Cards[cur].Rank, d.Cards[cur].Suit)
		row[i-1] = d.Cards[cur]
	}

	return row, nil
}
