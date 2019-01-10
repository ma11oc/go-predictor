package main

import (
	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var (
	origin = [52]uint8{
		3, 14, 25, 49, 18, 29, 40,
		7, 33, 44, 11, 22, 48, 2,
		13, 39, 6, 17, 28, 50, 21,
		32, 43, 10, 36, 47, 1, 27,
		38, 5, 16, 42, 9, 20, 31,
		51, 24, 35, 46, 15, 26, 37,
		4, 30, 41, 8, 19, 45, 12,
		23, 34, 52,
	}

	mm [90]*core.YearMatrix

	err error
)

func main() {
	od := core.NewOrderedDeck()
	om := core.NewOriginMatrix(&origin, od)
	// hm := core.NewHumansMatrix(om, od)
	// am := core.NewAngelsMatrix(om, od)

	// build all year matrices
	mm[0] = core.NewZeroYearMatrix(om)
	cur := mm[0]
	for i := 1; i < 90; i++ {
		if mm[i], err = cur.Next(om, od); err != nil {
			panic(err)
		}

		cur = mm[i]
	}

	/*
	 *     scs := spew.ConfigState{
	 *         Indent:           "  ",
	 *         ContinueOnMethod: true,
	 *         MaxDepth:         6,
	 *     }
	 *
	 *     scs.Dump(mm[89])
	 */

	for i := uint8(0); i < 90; i++ {
		for j := uint8(1); j <= 52; j++ {
			c, _ := core.NewCardFromNumber(j)
			mm[i].Decks.Main.GetVRow(c)
		}
	}
	// fmt.Println(mm[32].Decks.Main.AsNumbers())
	// mm[0].Decks.Drain.PrettyPrint()
}
