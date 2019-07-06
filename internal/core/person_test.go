package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

func calcBirthdayHelper(c *core.Card, age uint8) (time.Time, error) {
	bdays, _ := c.GetBirthdays()
	t := time.Now()
	y := t.Year() - int(age)

	for _, bday := range bdays {
		if t.YearDay() >= bday.YearDay() {
			return time.Date(y, bday.Month(), bday.Day(), 0, 0, 0, 0, time.UTC), nil
		}
	}

	return time.Date(y-1, bdays[0].Month(), bdays[0].Day(), 0, 0, 0, 0, time.UTC), nil
}

var _ = Describe("Internal/Core/Deck", func() {
	var (
		pc = &core.PersonConfig{}
	)

	Describe("resolve longterm card", func() {
		Context("when card is 3♦ and age is 40", func() {
			It("should return 9♣", func() {
				c, _ := core.NewCardFromNumber(29, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 40)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(22, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is J♣ and age is 37", func() {
			It("should return K♠", func() {
				c, _ := core.NewCardFromNumber(24, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 37)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(52, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 9♣ and age is 40", func() {
			It("should return J♦", func() {
				c, _ := core.NewCardFromNumber(22, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 40)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(37, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 6♦ and age is 66", func() {
			It("should return 10♠", func() {
				c, _ := core.NewCardFromNumber(32, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 66)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(49, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("resolve pluto/result card", func() {
		Context("when card is Q♦ and age is 31", func() {
			It("should return pluto as 2♦ and pluto result as 2♠", func() {
				c, _ := core.NewCardFromNumber(38, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 31)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(28, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(41, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is K♥ and age is 31", func() {
			It("should return pluto as 4♦ and pluto result as 5♠", func() {
				c, _ := core.NewCardFromNumber(13, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 31)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(30, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(44, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 10♥ and age is 31", func() {
			It("should return pluto as 2♥ and pluto result as 7♥", func() {
				c, _ := core.NewCardFromNumber(10, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 31)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(2, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(7, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 9♦ and age is 67", func() {
			It("should return pluto as 2♥ and pluto result as 7♥", func() {
				c, _ := core.NewCardFromNumber(35, locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 67)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(49, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(4, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
