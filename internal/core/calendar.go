package core

import "time"

// Day describes one day in calendar
type Day struct {
	Weekday time.Weekday
	Planet  *Planet
	Cards   struct {
		H *Card
		V *Card
	}
}

// Week describes a period within a year and contains a week
type Week struct {
	Days  []*Day
	Start time.Time
	End   time.Time
}

// Calendar contains one or two periods within a year
type Calendar struct {
	Weeks []*Week
}

// NewWeek returns a Week
func NewWeek(c *Card, ym *YearMatrix, pp *Planets, wd time.Weekday, b time.Time, e time.Time) (*Week, error) {
	var err error
	var hr, vr *Row
	var w *Week

	if hr, err = ComputeHRow(ym, c); err != nil {
		return nil, err
	}
	if vr, err = ComputeVRow(ym, c); err != nil {
		return nil, err
	}

	w = &Week{
		Start: b,
		End:   e,
	}

	days := []*Day{}

	for i := 0; i < 7; i++ {
		idx := int(wd) + i

		if idx >= 7 {
			idx -= 7
		}

		hc := hr[idx]
		vc := vr[idx]

		d := &Day{
			Weekday: time.Weekday(idx),
			Planet:  pp[idx],
			Cards: struct {
				H *Card
				V *Card
			}{
				hc,
				vc,
			},
		}

		days = append(days, d)
	}

	w.Days = days

	return w, nil
}
