package main

import (
	"bitbucket.org/shchukin_a/go-predictor/internal/core"
	"github.com/davecgh/go-spew/spew"
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
)

func main() {
	od := core.NewOrderedDeck()
	om := core.NewOriginMatrix(&origin, od)
	// hm := core.NewHumansMatrix(om, od)
	// am := core.NewAngelsMatrix(om, od)
	ym := core.NewZeroYearMatrix(om)

	scs := spew.ConfigState{
		Indent:           "  ",
		ContinueOnMethod: true,
		MaxDepth:         6,
	}

	scs.Dump(ym)
	scs.Dump(ym.Next(om, od))
}
