package core

// Matrix is a base primitive
type Matrix struct {
	Decks struct {
		Main   *Deck
		Drain  *Deck
		Source *Deck
	}
}

// Matrices is alias for all the matrices for [0, 89] years
type Matrices [90]*YearMatrix

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
	md, err := NewDeckFromSlice(*origin, od)
	if err != nil {
		panic(err)
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
		idx := om.Decks.Main.Cards[i].ID - 1
		sd.Cards[i] = od.Cards[om.Decks.Main.Cards[idx].ID-1]
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
		if dd.Cards[i], err = od.FindCardByNumber(idx + 1); err != nil {
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

// MatrixIterator describes Iterator for matrices
type MatrixIterator interface {
	Next(m *Matrix, d *Deck) (*Matrix, error)
}

// YearMatrix is a base primitive
type YearMatrix struct {
	Year uint8
	*Matrix
	MatrixIterator
}

// Next is the Matrix iterator
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
		if oIdx, err = om.Decks.Main.indexOf(m.Decks.Main.Cards[i].ID); err != nil {
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

// NewZeroYearMatrix receives origin Matrix and returns YearMatrix for 0 year
func NewZeroYearMatrix(om *Matrix) *YearMatrix {
	return &YearMatrix{
		Year:   0,
		Matrix: om,
	}
}

// NewMatrices receives origin Matrix and ordered Deck and
// returns all the matrices for years [0, 89]
func NewMatrices(om *Matrix, od *Deck) *Matrices {
	var err error

	mm := &Matrices{
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
