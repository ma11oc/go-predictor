package main

import (

	// "github.com/pkg/profile"

	"fmt"
	"time"

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

	od         *core.Deck
	om, hm, am *core.Matrix
	mm         [90]*core.YearMatrix

	err error

	d *core.Deck
)

func init() {
	// core.PrintAllPeriodicityCicles()

}

func main() {
	// CPU profiling by default
	// defer profile.Start().Stop()
	// Memory profiling
	// defer profile.Start(profile.MemProfile).Stop()

	// core.LoadLocales("locales/ru-RU.yaml", "locales/en-US.yaml")
	core.LoadLocales("locales/ru-RU.yaml")

	od := core.NewOrderedDeck("ru-RU")
	for i, v := range od.Cards {
		// scs.Dump(v)
		fmt.Printf("%2v: %v\n", i, v.Title)
		// scs.Dump(v)
		// fmt.Printf("%2v: %v\n", i, v.Meanings.General.Keywords)
		// fmt.Printf("%2v: %v\n", i, v.Meanings.Longterm.Keywords)
	}
	// od := core.NewOrderedDeck()
	om := core.NewOriginMatrix(&origin, od)
	hm := core.NewHumansMatrix(om, od)
	// am := core.NewAngelsMatrix(om, od)
	mm := core.NewBunchOfYearMatrices(om, od)
	b := time.Date(1986, time.April, 15, 0, 0, 0, 0, time.UTC)
	p, _ := core.NewPerson(b, od, mm, hm)

	scs := spew.ConfigState{
		Indent:   "  ",
		MaxDepth: 3,
	}
	scs.Dump(p)

	// fmt.Println(hm)
	// fmt.Println(am)
	// fmt.Println(mm)

}
