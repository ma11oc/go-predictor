package core

import (
	"fmt"
	"time"
)

// TODO: symbol is a rune
type planet struct {
	order  int
	name   string
	symbol string
}

type periodicityCircle struct {
	Start time.Time
	End   time.Time
}

// https://en.wikipedia.org/wiki/Astrological_symbols
// Transpluto \u2be8

var (
	Planets = [7]*planet{
		&planet{
			order:  1,
			name:   "mercury",
			symbol: "\u263f", // ☿
		},
		&planet{
			order:  2,
			name:   "venus",
			symbol: "\u2640", // ♀
		},
		&planet{
			order:  3,
			name:   "mars",
			symbol: "\u2642", // ♂
		},
		&planet{
			order:  4,
			name:   "jupiter",
			symbol: "\u2643", // ♃
		},
		&planet{
			order:  5,
			name:   "saturn",
			symbol: "\u2644", // ♄
		},
		&planet{
			order:  6,
			name:   "uranus",
			symbol: "\u2645", // ♅
		},
		&planet{
			order:  7,
			name:   "neptune",
			symbol: "\u2646", // ♆
		},
		/*
		 * &planet{
		 *     order:  8,
		 *     name:   "pluto",
		 *     symbol: "\u2647", // ♇
		 * },
		 */
	}
)

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func getAllPeriodicityCircles() [7][54]periodicityCircle {
	d := date(1999, 12, 31)
	m := [7][54]periodicityCircle{}

	for x := 0; x < 7; x++ {
		for y := 0; y < 54; y++ {
			m[x][y].Start = d.AddDate(0, 0, (x*52)+(y+1))
			if x != 6 {
				m[x][y].End = m[x][y].Start.AddDate(0, 0, 51)
			} else {
				m[x][y].End = m[x][y].Start.AddDate(0, 0, 53)
			}
		}
	}

	return m
}

func PrintAllPeriodicityCicles() {
	m := getAllPeriodicityCircles()

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

func GetCurrentPeriodicityCircles(t time.Time) [7]periodicityCircle {
	// just add 52 days from birthday 7 times
	c := getAllPeriodicityCircles()
	r := [7]periodicityCircle{}

	date := time.Date(2000, t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	days := int(date.Sub(time.Date(date.Year(), 1, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)

	x := int(days / 54)
	y := int(days+x*2) % 54

	for i := 0; i < 7; i++ {
		if x+i+1 >= 7 {
			x = (i + 1) * -1
		}

		r[i] = c[x+i+1][y]
	}

	return r
}
