package main

import (

	// "github.com/pkg/profile"

	"time"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/text/language"
)

func main() {
	// CPU profiling by default
	// defer profile.Start().Stop()
	// Memory profiling
	// defer profile.Start(profile.MemProfile).Stop()

	locales := core.BuildLocales("locales/ru-RU.yaml")
	lang := language.Make("ru-RU")

	pc := &core.PersonConfig{
		Name:     "Requester1",
		Birthday: time.Date(1986, time.April, 15, 0, 0, 0, 0, time.UTC),
		Gender:   core.Male,
		Environment: []*core.PersonConfig{
			&core.PersonConfig{
				Name:        "Requester2",
				Birthday:    time.Date(1986, time.July, 19, 0, 0, 0, 0, time.UTC),
				Gender:      core.Male,
				Environment: nil,
			},
			&core.PersonConfig{
				Name:        "Requester3",
				Birthday:    time.Date(2014, time.October, 11, 0, 0, 0, 0, time.UTC),
				Gender:      core.Male,
				Environment: nil,
			},
			&core.PersonConfig{
				Name:        "Requester4",
				Birthday:    time.Date(2014, time.December, 31, 0, 0, 0, 0, time.UTC),
				Gender:      core.Male,
				Environment: nil,
			},
		},
	}

	p, _ := core.NewPerson(pc, locales[lang])

	scs := spew.ConfigState{
		Indent:   "|---",
		MaxDepth: 5,
	}
	scs.Dump(p)

}
