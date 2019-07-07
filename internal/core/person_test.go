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
				c, _ := core.NewCardFromString("3♦", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 40)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("9♣", locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is J♣ and age is 37", func() {
			It("should return K♠", func() {
				c, _ := core.NewCardFromString("J♣", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 37)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("K♠", locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 9♣ and age is 40", func() {
			It("should return J♦", func() {
				c, _ := core.NewCardFromString("9♣", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 40)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("J♦", locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 6♦ and age is 66", func() {
			It("should return 10♠", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 66)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("10♠", locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("resolve pluto/result card", func() {
		Context("when card is Q♦ and age is 31", func() {
			It("should return pluto as 2♦ and pluto/result as 2♠", func() {
				c, _ := core.NewCardFromString("Q♦", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 31)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("2♦", locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromString("2♠", locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is K♥ and age is 31", func() {
			It("should return pluto as 4♦ and pluto/result as 5♠", func() {
				c, _ := core.NewCardFromString("K♥", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 31)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("4♦", locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromString("5♠", locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 10♥ and age is 31", func() {
			It("should return pluto as 2♥ and pluto/result as 7♥", func() {
				c, _ := core.NewCardFromString("10♥", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 31)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("2♥", locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromString("7♥", locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 9♦ and age is 67", func() {
			It("should return pluto as 10♠ and pluto/result as 4♥", func() {
				c, _ := core.NewCardFromString("9♦", locale)
				pc.Birthday, _ = calcBirthdayHelper(c, 67)

				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromString("10♠", locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromString("4♥", locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
