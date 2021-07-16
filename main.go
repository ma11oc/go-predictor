package main

import (

	// "github.com/pkg/profile"

	// "time"

	"time"

	"github.com/ma11oc/go-predictor/internal/core"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/text/language"
)

func main() {
	// CPU profiling by default
	// defer profile.Start().Stop()
	// Memory profiling
	// defer profile.Start(profile.MemProfile).Stop()

	locales := core.MustBuildLocales("locales/ru-RU.yaml")
	lang := language.Make("ru-RU")

	/*
	 * pc := &core.PersonProfile{
	 *     Name:     "Requester1",
	 *     Birthday: time.Date(1986, time.April, 15, 0, 0, 0, 0, time.UTC),
	 *     Gender:   core.Male,
	 * }
	 */

	// p, _ := core.NewPerson(pc, locales[lang])

	// cal, _ := core.ComputeCalendar(time.Date(1966, time.October, 24, 0, 0, 0, 0, time.UTC),
	cal, _ := core.ComputeCalendar(time.Date(1986, time.April, 15, 0, 0, 0, 0, time.UTC),
		locales[lang].GetOrderedDeck(), locales[lang].GetPlanets(), 2019, locales[lang].GetYearMatrices())

	scs := spew.ConfigState{
		Indent:   "    ",
		MaxDepth: 7,
	}
	scs.Dump(cal)

}
