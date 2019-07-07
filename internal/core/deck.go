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
// in the following order (index: id rank suit):
//    0:  1  A♥ | 13: 14  A♣ | 26: 27  A♦ | 39: 40  A♠
//    1:  2  2♥ | 14: 15  2♣ | 27: 28  2♦ | 40: 41  2♠
//    2:  3  3♥ | 15: 16  3♣ | 28: 29  3♦ | 41: 42  3♠
//    3:  4  4♥ | 16: 17  4♣ | 29: 30  4♦ | 42: 43  4♠
//    4:  5  5♥ | 17: 18  5♣ | 30: 31  5♦ | 43: 44  5♠
//    5:  6  6♥ | 18: 19  6♣ | 31: 32  6♦ | 44: 45  6♠
//    6:  7  7♥ | 19: 20  7♣ | 32: 33  7♦ | 45: 46  7♠
//    7:  8  8♥ | 20: 21  8♣ | 33: 34  8♦ | 46: 47  8♠
//    8:  9  9♥ | 21: 22  9♣ | 34: 35  9♦ | 47: 48  9♠
//    9: 10 10♥ | 22: 23 10♣ | 35: 36 10♦ | 48: 49 10♠
//   10: 11  J♥ | 23: 24  J♣ | 36: 37  J♦ | 49: 50  J♠
//   11: 12  Q♥ | 24: 25  Q♣ | 37: 38  Q♦ | 50: 51  Q♠
//   12: 13  K♥ | 25: 26  K♣ | 38: 39  K♦ | 51: 52  K♠
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
//
//     | Jan  Feb  Mar  Apr  May  Jun  Jul  Aug  Sep  Oct  Nov  Dec
//  ---+-----------------------------------------------------------
//   1 |  K♠   J♠   9♠   7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦
//   2 |  Q♠  10♠   8♠   6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦
//   3 |  J♠   9♠   7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦
//   4 | 10♠   8♠   6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦
//   5 |  9♠   7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣
//   6 |  8♠   6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣
//   7 |  7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣
//   8 |  6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣
//   9 |  5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣
//  10 |  4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣
//  11 |  3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣
//  12 |  2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣
//  13 |  A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣
//  14 |  K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣
//  15 |  Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣
//  16 |  J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣
//  17 | 10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣
//  18 |  9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥
//  19 |  8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥
//  20 |  7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥
//  21 |  6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥
//  22 |  5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥
//  23 |  4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥
//  24 |  3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥   7♥
//  25 |  2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥   6♥
//  26 |  A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥   7♥   5♥
//  27 |  K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥   6♥   4♥
//  28 |  Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥   7♥   5♥   3♥
//  29 |  J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥   6♥   4♥   2♥
//  30 | 10♣        6♣   4♣   2♣   K♥   J♥   9♥   7♥   5♥   3♥   A♥
//  31 |  9♣        5♣        A♣       10♥   8♥        4♥        🃏
func (d Deck) FindCardByBirthday(t time.Time) (*Card, error) {
	idx := 54 - (t.Day() + (int(t.Month()) * 2)) + 1

	if idx <= 0 || idx > 52 {
		return nil, fmt.Errorf("Got invalid card number `%v` for birthday: %v", idx, t)
	}

	return d.FindCardByNumber(uint8(idx))
}

// FindCardByIndex returns appropriate card from a deck
func (d Deck) FindCardByIndex(i uint8) (*Card, error) {
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

func (d Deck) CalcHRow(c *Card) ([7]*Card, error) {
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

// CalcVRow ...
// 6: [48, 41, 34, 27, 20, 13, 6]      [-1, -8,  -15, ...]
// 5: [47, 40, 33, 26, 19, 12, 5]      [-2, -9,  -16, ...]
// 4: [46, 39, 32, 25, 18, 11, 4, 51]  [-3, -10, -17, ...]
// 3: [45, 38, 31, 24, 17, 10, 3, 50]  [-4, -11, -18, ...]
// 2: [44, 37, 30, 23, 16, 9,  2, 49]  [-5, -12, -19, ...]
// 1: [43, 36, 29, 22, 15, 8,  1]      [-6, -13, -20, ...]
// 0: [42, 35, 28, 21, 14, 7,  0]      [-7, -14, -21, ...]
func (d Deck) CalcVRow(c *Card) ([7]*Card, error) {
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
