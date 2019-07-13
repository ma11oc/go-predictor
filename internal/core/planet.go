package core

import (
	"fmt"
	"time"
)

var (
	// PlanetsOrder is the order of planet and must be const
	PlanetsOrder = [7]string{"mercury", "venus", "mars", "jupiter", "saturn", "uranus", "neptune"}
)

// Planet represents a planet primitive
type Planet struct {
	ID     uint8  `yaml:"id"      validate:"min=1,max=7"`
	Name   string `yaml:"name"    validate:"nonzero"`
	Symbol string `yaml:"symbol"  validate:"nonzero,regexp=^(☿|♀|♂|♃|♄|♅|♆)$"`
}

// Planets is an alias for map of planets
type Planets [7]*Planet

// Cycle has start and end dates
type Cycle struct {
	Start time.Time
	End   time.Time
}

// Cycles is alias for cycles matrix
type Cycles [7][54]*Cycle

// PlanetCycle represents a planet cycle primitive
type PlanetCycle struct {
	Cards struct {
		H *Card // Card from Horizontal Row
		V *Card // Card from Vertical Row
	}
	Planet *Planet
	Start  time.Time
	End    time.Time
}

// PlanetCycles is alias to array of PlanetCycle
type PlanetCycles [7]*PlanetCycle

// Row is alias, represents horizontal or vertical row in matrix
type Row [7]*Card

// Rows is alias for set of horizontal and vertical row
type Rows struct {
	H *Row `yaml:"hrow" validate:"nonzero,min=7,max=7"`
	V *Row `yaml:"vrow" validate:"nonzero,min=6,max=7"`
}

// NewCyclesTable returns matrix with planet cycles during a year
// See README.md > Appendix > Planet Cycles
func NewCyclesTable() *Cycles {
	d := time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
	m := Cycles{}

	for x := 0; x < 7; x++ {
		for y := 0; y < 54; y++ {
			m[x][y] = &Cycle{
				Start: d.AddDate(0, 0, (x*52)+(y+1)),
			}

			if x != 6 {
				m[x][y].End = m[x][y].Start.AddDate(0, 0, 51)
			} else {
				m[x][y].End = m[x][y].Start.AddDate(0, 0, 53)
			}
		}
	}

	return &m
}

// PrintCycles prints the table of cycles to stdout
func PrintCycles() {
	m := NewCyclesTable()

	for y := 0; y < 54; y++ {
		for x := 0; x < 7; x++ {
			fmt.Printf(" %v-%v |",
				m[x][y].Start.Format("02/01"),
				m[x][y].End.Format("02/01"),
			)
		}
		fmt.Println()
	}
}
