package core

// Matrix is a base struct.
type Matrix struct {
	Decks struct {
		Main   *Deck
		Drain  *Deck
		Source *Deck
	}
}

/*
 * func (m Matrix) PrettyPrint() {
 *     for i := 0; i < 52; i++ {
 *         fmt.Printf("%4v", )
 *     }
 * }
 */

// NewOriginMatrix returns filled in OriginMatrix.
// Accepts two arguments:
//   - origin slice
//   - ordered deck
// where:
//   main[i] = origin[i]
//   drain[i] = ordered_deck[i]
//   source[i] = origin[i]
func NewOriginMatrix(origin *[52]uint8, od *Deck) *Matrix {
	var c *Card
	var err error

	md := &Deck{}

	for i := uint8(0); i < 52; i++ {
		if c, err = NewCardFromNumber(origin[i]); err != nil {
			panic("Unable to build OriginMatrix")
		}
		md.Cards[i] = c
	}

	return &Matrix{
		Decks: struct {
			Main   *Deck
			Drain  *Deck
			Source *Deck
		}{
			md,
			od,
			md,
		},
	}

}

// NewHumansMatrix returns appropriate matrix.
// It accepts two args:
//   - built origin matrix
//   - ordered deck
// where:
//   main[i] = origin[i-1];
//   drain[i] = i;
//   source[i] = origin[origin[i-1]-1];
func NewHumansMatrix(om *Matrix, od *Deck) *Matrix {
	sd := &Deck{
		Cards: [52]*Card{},
	}

	for i := uint8(0); i < 52; i++ {
		idx := om.Decks.Main.Cards[i].Number - 1
		sd.Cards[i] = od.Cards[om.Decks.Main.Cards[idx].Number-1]
	}

	return &Matrix{
		Decks: struct {
			Main   *Deck
			Drain  *Deck
			Source *Deck
		}{
			om.Decks.Main,
			od, // ordered deck
			sd,
		},
	}

}

// NewAngelsMatrix return appropriate matrix.
// It accepts two args:
//   - built origin matrix
//   - ordered deck
// where:
//   main[i] = ordered_deck[i];
//   drain[i] = origin.indexof(i)+1;
//   source[i] = origin[i-1];
func NewAngelsMatrix(om *Matrix, od *Deck) *Matrix {
	var idx uint8
	var err error

	dd := &Deck{
		Cards: [52]*Card{},
	}

	for i := uint8(0); i < 52; i++ {
		if idx, err = om.Decks.Main.indexOf(i + 1); err != nil {
			panic(err)
		}
		if dd.Cards[i], err = od.GetCardByNumber(idx + 1); err != nil {
			panic(err)
		}
	}

	return &Matrix{
		Decks: struct {
			Main   *Deck
			Drain  *Deck
			Source *Deck
		}{
			od, // ordered deck
			dd,
			om.Decks.Main,
		},
	}

}

/*
 * year
 *  m  main = year_matrix[i-1];
 *  d  drain = i;
 *  s  source = origin[i-1];
 * }
 */

type MatrixIterator interface {
	Next(m *Matrix, d *Deck) (*Matrix, error)
}

type YearMatrix struct {
	Year uint8
	*Matrix
	MatrixIterator
}

func (m YearMatrix) Next(om *Matrix, od *Deck) (*YearMatrix, error) {
	var oIdx uint8
	var mIdx uint8
	var err error

	next := &YearMatrix{
		Matrix: &Matrix{
			Decks: struct {
				Main   *Deck
				Drain  *Deck
				Source *Deck
			}{
				&Deck{},
				od,
				om.Decks.Main,
			},
		},
	}

	for i := uint8(0); i < 52; i++ {
		if oIdx, err = om.Decks.Main.indexOf(m.Decks.Main.Cards[i].Number); err != nil {
			panic(err)
		}
		if mIdx, err = m.Decks.Main.indexOf(oIdx + 1); err != nil {
			panic(err)
		}

		next.Decks.Main.Cards[mIdx] = m.Decks.Main.Cards[i]
		next.Year = m.Year + 1
	}

	return next, nil
}

func NewZeroYearMatrix(om *Matrix) *YearMatrix {
	return &YearMatrix{
		Year:   0,
		Matrix: om,
	}
}

func NewBunchOfYearMatrices(om *Matrix, od *Deck) [90]*YearMatrix {
	var err error

	mm := [90]*YearMatrix{
		&YearMatrix{
			Year:   0,
			Matrix: om,
		},
	}

	cur := mm[0]
	for i := 1; i < 90; i++ {
		if mm[i], err = cur.Next(om, od); err != nil {
			panic(err)
		}

		cur = mm[i]
	}

	return mm
}

/*
 *
 * func getMatrixByYear(y uint8) (*matrix, error) {
 *     var i uint8
 *
 *     if y == 0 {
 *         return nil, fmt.Errorf("Invalid year number")
 *     }
 *
 *     if y < 90 {
 *         i = y
 *     } else {
 *         i = y % 90
 *     }
 *
 *     return matrices[i-1], nil
 * }
 */
