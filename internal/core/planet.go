package core

import (
	"time"
)

// Planet represents a planet primitive
type Planet struct {
	ID     uint8  `yaml:"id"      validate:"min=1,max=7"`
	Name   string `yaml:"name"    validate:"nonzero"`
	Symbol string `yaml:"symbol"  validate:"nonzero,regexp=^(☿|♀|♂|♃|♄|♅|♆)$"`
}

// PlanetCycle represents a planet cycle primitive
type PlanetCycle struct {
	Card   *Card
	Planet *Planet
	Start  time.Time
	End    time.Time
}

// NewBunchOfPlanetCycles returns matrix with planet cycles during a year
// See README.md > Appendix > Planet Cycles
func NewBunchOfPlanetCycles() *[7][54]*PlanetCycle {
	d := time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
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
 *     m := NewBunchOfPlanetCycles()
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

// GetCurrentPlanetCycles calculates planet cycles according to birthday
// and returns array with them
func GetCurrentPlanetCycles(t time.Time, pc *[7][54]*PlanetCycle, planets *[7]*Planet) *[7]*PlanetCycle {
	if pc == nil || planets == nil {
		return nil
	}

	r := [7]*PlanetCycle{}

	date := time.Date(2000, t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	days := int(date.Sub(time.Date(date.Year(), 1, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)

	x := int(days / 54)
	y := int(days+x*2) % 54

	// just add 52 days from birthday 7 times
	for i := 0; i < 7; i++ {
		if x+i+1 >= 7 {
			x = (i + 1) * -1
		}

		r[i] = (*pc)[x+i+1][y]
		r[i].Planet = planets[i]
	}

	return &r
}
