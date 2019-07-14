package core

import (
	"fmt"
	"time"
)

// ComputeMainCards receives birthday, ordered deck and humans matrix and
// returns 3 cards: Main, Drain and Source or error
func ComputeMainCards(b time.Time, od *Deck, hm *Matrix) (*Card, *Card, *Card, error) {
	var err error
	var idx uint8

	var mc, dc, sc *Card

	if mc, err = od.FindCardByBirthday(b); err != nil {
		return nil, nil, nil, err
	}

	if idx, err = hm.Decks.Main.indexOf(mc.ID); err != nil {
		return nil, nil, nil, err
	}

	if dc, err = hm.Decks.Drain.FindCardByIndex(idx); err != nil {
		return nil, nil, nil, err
	}

	if sc, err = hm.Decks.Source.FindCardByIndex(idx); err != nil {
		return nil, nil, nil, err
	}

	return mc, dc, sc, nil
}

// ComputeLongtermCard receives array of year matrices, age and Main card of Person and
// returns appropriate longterm card or error
func ComputeLongtermCard(mm *Matrices, c *Card, age uint8) (*Card, error) {
	var idx uint8
	var err error

	var lc *Card

	ym := mm[age/7]

	if idx, err = ym.Matrix.Decks.Main.indexOf(c.ID); err != nil {
		return nil, err
	}

	idx += age%7 + 1
	if idx >= 52 {
		idx -= 52
	}

	if lc, err = ym.Matrix.Decks.Main.FindCardByIndex(idx); err != nil {
		return nil, err
	}

	return lc, nil
}

// ComputePlutoCards receives YearMatrix and Main card of Person and returns
// pluto card and pluto result card or error
func ComputePlutoCards(m *YearMatrix, c *Card) (*Card, *Card, error) {
	var idx uint8
	var err error
	var pc, rc *Card

	if idx, err = m.Decks.Main.indexOf(c.ID); err != nil {
		return nil, nil, err
	}

	if idx+8 >= 52 {
		idx = idx - 52
	}

	idx += 8

	if pc, err = m.Decks.Main.FindCardByIndex(idx); err != nil {
		return nil, nil, err
	}

	if idx+1 >= 52 {
		idx = idx - 52
	}

	idx++

	if rc, err = m.Decks.Main.FindCardByIndex(idx); err != nil {
		return nil, nil, err
	}

	return pc, rc, nil
}

// ComputeHRow receives year matrix and Main card of a Person and
// retuns Row (array with 7 cards)
func ComputeHRow(m *YearMatrix, c *Card) (*Row, error) {
	var row *Row
	var ci uint8 // card index
	var err error

	row = &Row{}

	// card is a starting point to count from
	if ci, err = m.Decks.Main.indexOf(c.ID); err != nil {
		return nil, err
	}

	for i := uint8(1); i <= 7; i++ {
		cur := ci + i
		if cur >= 52 {
			cur -= 52
		}

		row[i-1] = m.Decks.Main.Cards[cur]
	}

	return row, nil
}

// ComputeVRow receives year matrix and Main card of a Person and
// retuns Row (array with 7 cards)
// 6: [48, 41, 34, 27, 20, 13, 6]      [-1, -8,  -15, ...]
// 5: [47, 40, 33, 26, 19, 12, 5]      [-2, -9,  -16, ...]
// 4: [46, 39, 32, 25, 18, 11, 4, 51]  [-3, -10, -17, ...]
// 3: [45, 38, 31, 24, 17, 10, 3, 50]  [-4, -11, -18, ...]
// 2: [44, 37, 30, 23, 16, 9,  2, 49]  [-5, -12, -19, ...]
// 1: [43, 36, 29, 22, 15, 8,  1]      [-6, -13, -20, ...]
// 0: [42, 35, 28, 21, 14, 7,  0]      [-7, -14, -21, ...]
func ComputeVRow(m *YearMatrix, c *Card) (*Row, error) {
	var row *Row
	var ci uint8 // card index
	var err error

	row = &Row{}

	if ci, err = m.Decks.Main.indexOf(c.ID); err != nil {
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
		// fmt.Printf("%v: %2v | %v%v\n", i, m.Decks.Main.Cards[cur].ID, m.Decks.Main.Cards[cur].Rank, m.Decks.Main.Cards[cur].Suit)
		row[i-1] = m.Decks.Main.Cards[cur]
	}

	return row, nil
}

// ComputePlanetCycles calculates planet cycles according to birthday
// and returns array with them
func ComputePlanetCycles(b time.Time, cc *Cycles, pp *Planets, hr *Row, vr *Row) (*PlanetCycles, error) {
	// var err error

	pcc := &PlanetCycles{}

	// since table based on leap year, we need to move birthday (b)
	// to leap year, 2000, for instance
	date := time.Date(2000, b.Month(), b.Day(), 0, 0, 0, 0, time.UTC)

	// find coords in PlanetCyclesMatrix to start count
	x := int(date.YearDay() / 54)
	y := int(date.YearDay()+x*2) % 54

	for i := 0; i < 7; i++ {
		// if we reached the end of a row in cycles matrix,
		// start from the beginning
		if x+i+1 >= 7 {
			x = (i + 1) * -1
		}

		c := (*cc)[x+i+1][y]

		pcc[i] = &PlanetCycle{
			Start:  c.Start,
			End:    c.End,
			Planet: pp[i],
			Cards: struct {
				H *Card
				V *Card
			}{
				hr[i],
				vr[i],
			},
		}

	}

	return pcc, nil
}

// ComputePersonalCards receives Main card, gender, age, features and Locale and
// returns PersonalCards ([3]*Card) or error
// Men:
//   - in spite of age, each man has Jack with the same Suit as his main card,
//     except the case when a man already has Jack with the same Suit
//     as a main card
//   - if a man over 36 years old, he has King with the same Suit
//   - if a man is a business owner with at least 2 employees and more,
//     or a man is a chief (behaves like a chief) he has King with the same Suit
// Women:
//   - women may have up to 3 personal cards at the same time
//   - in spite of age, each woman has Queen with the same Suit as her main card
//     except the case when a woman already has Queen with the same Suit
//     as a main card
//   - if a woman is an actress, a writer or an artist, she has the Jack with
//     the same Suit
//   - if a woman is a business owner with at least 2 employees and more,
//     or a woman is a chief (behaves like a chief) she has King with the same Suit
//   - women younger than 20 years old has Jack with the same Suit
func ComputePersonalCards(c *Card, g Gender, f Feature, a uint8, l *Locale) (*PersonalCards, error) {
	var err error
	var card *Card

	pcc := PersonalCards{}

	switch g {
	case Male:
		if c.Rank != "J" {
			if card, err = l.FindCardByString("J" + c.Suit); err != nil {
				return nil, err
			}
			pcc = append(pcc, card)
		}

		if a >= 36 || f&Business > 0 {
			if card, err = l.FindCardByString("K" + c.Suit); err != nil {
				return nil, err
			}
			pcc = append(pcc, card)
		}

	case Female:
		if a <= 20 || f&Creator > 0 {
			if card, err = l.FindCardByString("J" + c.Suit); err != nil {
				return nil, err
			}
			pcc = append(pcc, card)
		}

		if c.Rank != "Q" {
			if card, err = l.FindCardByString("Q" + c.Suit); err != nil {
				return nil, err
			}
			pcc = append(pcc, card)
		}

		if f&Business > 0 {
			if card, err = l.FindCardByString("K" + c.Suit); err != nil {
				return nil, err
			}
			pcc = append(pcc, card)
		}

	default:
		return nil, fmt.Errorf("Unable to compute personal cards: no gender has been specified")
	}

	return &pcc, nil
}
