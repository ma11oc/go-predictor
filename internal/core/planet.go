package core

import (
	"time"
)

// TODO: Symbol is a rune
type Planet struct {
	ID     uint8
	Name   string
	Symbol string
}

type PlanetCycle struct {
	Card   *Card
	Planet *Planet
	Start  time.Time
	End    time.Time
}

// https://en.wikipedia.org/wiki/Astrological_symbols
// Transpluto \u2be8

var (
	Planets = [7]*Planet{
		&Planet{
			ID:     1,
			Name:   "mercury",
			Symbol: "\u263f", // ☿
		},
		&Planet{
			ID:     2,
			Name:   "venus",
			Symbol: "\u2640", // ♀
		},
		&Planet{
			ID:     3,
			Name:   "mars",
			Symbol: "\u2642", // ♂
		},
		&Planet{
			ID:     4,
			Name:   "jupiter",
			Symbol: "\u2643", // ♃
		},
		&Planet{
			ID:     5,
			Name:   "saturn",
			Symbol: "\u2644", // ♄
		},
		&Planet{
			ID:     6,
			Name:   "uranus",
			Symbol: "\u2645", // ♅
		},
		&Planet{
			ID:     7,
			Name:   "neptune",
			Symbol: "\u2646", // ♆
		},
		/*
		 * &Planet{
		 *     ID:  8,
		 *     Name:   "pluto",
		 *     Symbol: "\u2647", // ♇
		 * },
		 */
	}
)

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func NewBunchOfPlanetCycles() *[7][54]*PlanetCycle {
	d := date(1999, 12, 31)
	m := [7][54]*PlanetCycle{}

	for x := 0; x < 7; x++ {
		for y := 0; y < 54; y++ {
			m[x][y] = &PlanetCycle{
				Start: d.AddDate(0, 0, (x*52)+(y+1)),
			}
			// m[x][y].Start = d.AddDate(0, 0, (x*52)+(y+1))
			if x != 6 {
				m[x][y].End = m[x][y].Start.AddDate(0, 0, 51)
			} else {
				m[x][y].End = m[x][y].Start.AddDate(0, 0, 53)
			}
		}
	}

	return &m
}

/*
 * func PrintAllPeriodicityCicles() {
 *     m := getAllPlanetCycles()
 *
 *     for y := 0; y < 54; y++ {
 *         for x := 0; x < 7; x++ {
 *             fmt.Printf(" %v-%v |",
 *                 m[x][y].Start.Format("02/01"),
 *                 m[x][y].End.Format("02/01"),
 *             )
 *         }
 *         fmt.Println()
 *     }
 * }
 */

func GetCurrentPlanetCycles(t time.Time, pc *[7][54]*PlanetCycle) *[7]*PlanetCycle {
	// just add 52 days from birthday 7 times

	if pc == nil {
		return nil
	}

	r := [7]*PlanetCycle{}

	date := time.Date(2000, t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	days := int(date.Sub(time.Date(date.Year(), 1, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)

	x := int(days / 54)
	y := int(days+x*2) % 54

	for i := 0; i < 7; i++ {
		if x+i+1 >= 7 {
			x = (i + 1) * -1
		}

		r[i] = (*pc)[x+i+1][y]
		r[i].Planet = Planets[i]
	}

	return &r
}
