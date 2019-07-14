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

var _ = Describe("internal/core/core", func() {

	var (
		conf = &core.PersonProfile{}
	)

	Describe("ComputeMainCards", func() {
		Context("when birthday is 1 of January", func() {
			It("should return (Main, Drain, Source) == (K♠, K♠, K♠)", func() {
				b := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
				mc, dc, sc, err := core.ComputeMainCards(b, od, hm)

				want, _ := core.NewCardFromString("K♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(mc).To(Equal(want))
				Expect(dc).To(Equal(want))
				Expect(sc).To(Equal(want))
			})

			It("should return (Main, Drain, Source) == (6♦, 9♣, 3♠)", func() {
				b := time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				mc, dc, sc, err := core.ComputeMainCards(b, od, hm)

				mcw, _ := core.NewCardFromString("6♦", locale)
				dcw, _ := core.NewCardFromString("9♣", locale)
				scw, _ := core.NewCardFromString("3♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(mc).To(Equal(mcw))
				Expect(dc).To(Equal(dcw))
				Expect(sc).To(Equal(scw))
			})
		})
	})

	Describe("ComputeLongtermCard", func() {
		Context("when card is 3♦ and age is 40", func() {
			It("should return 9♣", func() {
				c, _ := core.NewCardFromString("3♦", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 40)
				lc, err := core.ComputeLongtermCard(mm, c, 40)

				lcw, _ := core.NewCardFromString("9♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
		Context("when card is J♣ and age is 37", func() {
			It("should return K♠", func() {
				c, _ := core.NewCardFromString("J♣", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 37)
				lc, err := core.ComputeLongtermCard(mm, c, 37)

				lcw, _ := core.NewCardFromString("K♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
		Context("when card is 9♣ and age is 40", func() {
			It("should return J♦", func() {
				c, _ := core.NewCardFromString("9♣", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 40)
				lc, err := core.ComputeLongtermCard(mm, c, 40)

				lcw, _ := core.NewCardFromString("J♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
		Context("when card is 6♦ and age is 66", func() {
			It("should return 10♠", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 66)
				lc, err := core.ComputeLongtermCard(mm, c, 66)

				lcw, _ := core.NewCardFromString("10♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
	})

	Describe("ComputePlutoCards", func() {
		Context("when card is Q♦ and age is 31", func() {
			It("should return pluto as 2♦ and pluto/result as 2♠", func() {
				c, _ := core.NewCardFromString("Q♦", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 31)
				pc, rc, err := core.ComputePlutoCards(mm[31], c)

				pcw, _ := core.NewCardFromString("2♦", locale)
				rcw, _ := core.NewCardFromString("2♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(pc).To(Equal(pcw))
				Expect(rc).To(Equal(rcw))
			})
		})
		Context("when card is K♥ and age is 31", func() {
			It("should return pluto as 4♦ and pluto/result as 5♠", func() {
				c, _ := core.NewCardFromString("K♥", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 31)
				pc, rc, err := core.ComputePlutoCards(mm[31], c)

				pcw, _ := core.NewCardFromString("4♦", locale)
				rcw, _ := core.NewCardFromString("5♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(pc).To(Equal(pcw))
				Expect(rc).To(Equal(rcw))
			})
		})
		Context("when card is 10♥ and age is 31", func() {
			It("should return pluto as 2♥ and pluto/result as 7♥", func() {
				c, _ := core.NewCardFromString("10♥", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 31)
				pc, rc, err := core.ComputePlutoCards(mm[31], c)

				pcw, _ := core.NewCardFromString("2♥", locale)
				rcw, _ := core.NewCardFromString("7♥", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(pc).To(Equal(pcw))
				Expect(rc).To(Equal(rcw))
			})
		})
		Context("when card is 9♦ and age is 67", func() {
			It("should return pluto as 10♠ and pluto/result as 4♥", func() {
				c, _ := core.NewCardFromString("9♦", locale)
				conf.Birthday, _ = calcBirthdayHelper(c, 67)
				pc, rc, err := core.ComputePlutoCards(mm[67], c)

				pcw, _ := core.NewCardFromString("10♠", locale)
				rcw, _ := core.NewCardFromString("4♥", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(pc).To(Equal(pcw))
				Expect(rc).To(Equal(rcw))
			})
		})
	})

	Describe("ComputeHRow", func() {
		Context("when main card is 3♦", func() {
			It("should return a valid slice of cards", func() {
				card, _ := core.NewCardFromString("3♦", locale)
				row, err := core.ComputeHRow(mm[0], card)

				s := [7]string{"A♠", "7♥", "7♦", "5♠", "J♥", "9♣", "9♠"}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromString(v, locale)
					Expect(c).To(Equal(row[i]))
				}
			})
		})

		Context("when main card is 6♠", func() {
			It("should return a valid slice of cards", func() {
				card, _ := core.NewCardFromString("6♠", locale)
				row, err := core.ComputeHRow(mm[0], card)

				s := [7]string{"Q♥", "10♣", "8♦", "K♠", "3♥", "A♣", "Q♣"}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromString(v, locale)
					Expect(c).To(Equal(row[i]))
				}
			})
		})

		Context("when main card is 8♦", func() {
			It("should return a valid slice of cards", func() {
				card, _ := core.NewCardFromString("8♦", locale)
				row, err := core.ComputeHRow(mm[0], card)

				s := [7]string{"K♠", "3♥", "A♣", "Q♣", "10♠", "5♣", "3♦"}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromString(v, locale)
					Expect(c).To(Equal(row[i]))
				}
			})
		})

		Context("when main card is 4♣", func() {
			It("should return a valid slice of cards", func() {
				card, _ := core.NewCardFromString("4♣", locale)
				row, err := core.ComputeHRow(mm[0], card)

				s := [7]string{"2♦", "J♠", "8♣", "6♦", "4♠", "10♥", "10♦"}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromString(v, locale)
					Expect(c).To(Equal(row[i]))
				}
			})
		})

		Context("for all the cards in all the years", func() {
			It("should not raise an error", func() {
				for i := uint8(0); i < 90; i++ {
					for j := uint8(1); j <= 52; j++ {
						c, _ := core.NewCardFromNumber(j, locale)
						row, err := core.ComputeHRow(mm[i], c)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(row).Should(BeAssignableToTypeOf(&core.Row{}))
					}
				}
			})
		})
	})

	Describe("ComputeVRow", func() {
		Context("for 3♦", func() {
			It("should return a valid slice of cards", func() {
				card, _ := core.NewCardFromString("3♦", locale)
				row, err := core.ComputeVRow(mm[0], card)

				s := [7]string{"6♠", "K♣", "7♣", "A♥", "J♠", "9♠"}

				Expect(err).ShouldNot(HaveOccurred())
				Expect(row[6]).To(BeNil())

				for i, v := range s {
					c, _ := core.NewCardFromString(v, locale)
					Expect(row[i]).To(Equal(c))
				}
			})
		})

		Context("for 4♣", func() {
			It("should return a valid slice of cards", func() {
				card, _ := core.NewCardFromString("4♣", locale)
				row, err := core.ComputeVRow(mm[0], card)

				s := [7]string{"J♥", "10♠", "8♦", "8♥", "7♠", "3♠", "10♦"}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromString(v, locale)
					Expect(row[i]).To(Equal(c))
				}
			})
		})

		Context("for all the cards in all the years", func() {
			It("should not raise an error", func() {
				for i := uint8(0); i < 90; i++ {
					for j := uint8(1); j <= 52; j++ {
						c, _ := core.NewCardFromNumber(j, locale)
						row, err := core.ComputeVRow(mm[i], c)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(row).Should(BeAssignableToTypeOf(&core.Row{}))
					}
				}
			})
		})
	})

	Describe("ComputePersonalCards", func() {
		Context("for J♣, male, 28 years old, student", func() {
			It("should return empty array of personal cards", func() {
				c, _ := core.NewCardFromString("J♣", locale)
				pcc, err := core.ComputePersonalCards(c, core.Male, core.Feature(0x0), 28, locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(len((*pcc))).Should(Equal(0))
			})
		})

		Context("for Q♦, male, 40 years old", func() {
			It("should return array with 2 personal cards", func() {
				c, _ := core.NewCardFromString("Q♦", locale)
				pcc, err := core.ComputePersonalCards(c, core.Male, core.Feature(0x0), 40, locale)

				pc0, _ := core.NewCardFromString("J♦", locale)
				pc1, _ := core.NewCardFromString("K♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(*pcc)).Should(Equal(2))
				Expect(pc0.In((*pcc))).Should(Equal(true))
				Expect(pc1.In((*pcc))).Should(Equal(true))

			})
		})

		Context("for K♣, female, 40 years old", func() {
			It("should return array with 1 personal card", func() {
				c, _ := core.NewCardFromString("K♣", locale)
				pcc, err := core.ComputePersonalCards(c, core.Female, core.Feature(0x0), 40, locale)

				pc0, _ := core.NewCardFromString("Q♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(*pcc)).Should(Equal(1))
				Expect(pc0.In((*pcc))).Should(Equal(true))

			})
		})

		Context("for 3♦, male, 20 years old, businessman", func() {
			It("should return array with 2 personal cards", func() {
				c, _ := core.NewCardFromString("3♦", locale)
				pcc, err := core.ComputePersonalCards(c, core.Male, core.Business, 20, locale)

				pc0, _ := core.NewCardFromString("J♦", locale)
				pc1, _ := core.NewCardFromString("K♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(*pcc)).Should(Equal(2))
				Expect(pc0.In((*pcc))).Should(Equal(true))
				Expect(pc1.In((*pcc))).Should(Equal(true))
			})
		})
		Context("for 3♦, female, 30 years old, businesswoman and actress", func() {
			It("should return array with 3 personal cards", func() {
				c, _ := core.NewCardFromString("3♦", locale)
				pcc, err := core.ComputePersonalCards(c, core.Female, core.Business|core.Creator, 30, locale)

				pc0, _ := core.NewCardFromString("J♦", locale)
				pc1, _ := core.NewCardFromString("Q♦", locale)
				pc2, _ := core.NewCardFromString("K♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(*pcc)).Should(Equal(3))
				Expect(pc0.In((*pcc))).Should(Equal(true))
				Expect(pc1.In((*pcc))).Should(Equal(true))
				Expect(pc2.In((*pcc))).Should(Equal(true))
			})
		})
	})
})
