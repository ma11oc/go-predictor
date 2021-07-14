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
		pp = &core.PersonProfile{}
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
				pp.Birthday, _ = calcBirthdayHelper(c, 40)
				lc, err := core.ComputeLongtermCard(mm, c, 40)

				lcw, _ := core.NewCardFromString("9♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
		Context("when card is J♣ and age is 37", func() {
			It("should return K♠", func() {
				c, _ := core.NewCardFromString("J♣", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 37)
				lc, err := core.ComputeLongtermCard(mm, c, 37)

				lcw, _ := core.NewCardFromString("K♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
		Context("when card is 9♣ and age is 40", func() {
			It("should return J♦", func() {
				c, _ := core.NewCardFromString("9♣", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 40)
				lc, err := core.ComputeLongtermCard(mm, c, 40)

				lcw, _ := core.NewCardFromString("J♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
		})
		Context("when card is 6♦", func() {
			It("should return 6♣, when age is 32", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 32)
				lc, err := core.ComputeLongtermCard(mm, c, 32)

				lcw, _ := core.NewCardFromString("6♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
			It("should return J♦, when age is 33", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 33)
				lc, err := core.ComputeLongtermCard(mm, c, 33)

				lcw, _ := core.NewCardFromString("J♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
			It("should return Q♥, when age is 34", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 34)
				lc, err := core.ComputeLongtermCard(mm, c, 34)

				lcw, _ := core.NewCardFromString("Q♥", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
			It("should return 4♣, when age is 35", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 35)
				lc, err := core.ComputeLongtermCard(mm, c, 35)

				lcw, _ := core.NewCardFromString("4♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(lc).To(Equal(lcw))
			})
			It("should return 10♠, when age is 66", func() {
				c, _ := core.NewCardFromString("6♦", locale)
				pp.Birthday, _ = calcBirthdayHelper(c, 66)
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
				pp.Birthday, _ = calcBirthdayHelper(c, 31)
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
				pp.Birthday, _ = calcBirthdayHelper(c, 31)
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
				pp.Birthday, _ = calcBirthdayHelper(c, 31)
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
				pp.Birthday, _ = calcBirthdayHelper(c, 67)
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

	Describe("ComputeKarmaCards", func() {
		Context("for 4♥", func() {
			It("should return array with 2 cards", func() {
				c, _ := core.NewCardFromString("4♥", locale)
				kcc, err := core.ComputeKarmaCards(c, hm, locale)

				kc1, _ := core.NewCardFromString("4♦", locale)
				kc2, _ := core.NewCardFromString("10♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(kc1.In((*kcc))).Should(Equal(true))
				Expect(kc2.In((*kcc))).Should(Equal(true))
				Expect(len((*kcc))).Should(Equal(2))
			})
		})
		Context("for J♥", func() {
			It("should return array with 2 cards", func() {
				c, _ := core.NewCardFromString("J♥", locale)
				kcc, err := core.ComputeKarmaCards(c, hm, locale)

				kc1, _ := core.NewCardFromString("K♠", locale)
				kc2, _ := core.NewCardFromString("8♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(kc1.In((*kcc))).Should(Equal(true))
				Expect(kc2.In((*kcc))).Should(Equal(true))
				Expect(len((*kcc))).Should(Equal(2))
			})
		})
		Context("for 8♣", func() {
			It("should return array with 2 cards", func() {
				c, _ := core.NewCardFromString("8♣", locale)
				kcc, err := core.ComputeKarmaCards(c, hm, locale)

				kc1, _ := core.NewCardFromString("J♥", locale)
				kc2, _ := core.NewCardFromString("K♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(kc1.In((*kcc))).Should(Equal(true))
				Expect(kc2.In((*kcc))).Should(Equal(true))
				Expect(len((*kcc))).Should(Equal(2))
			})
		})
		Context("for K♠", func() {
			It("should return array with 2 cards", func() {
				c, _ := core.NewCardFromString("K♠", locale)
				kcc, err := core.ComputeKarmaCards(c, hm, locale)

				kc1, _ := core.NewCardFromString("8♣", locale)
				kc2, _ := core.NewCardFromString("J♥", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(kc1.In((*kcc))).Should(Equal(true))
				Expect(kc2.In((*kcc))).Should(Equal(true))
				Expect(len((*kcc))).Should(Equal(2))
			})
		})
		Context("for 2♥", func() {
			It("should return array with 2 cards", func() {
				c, _ := core.NewCardFromString("2♥", locale)
				kcc, err := core.ComputeKarmaCards(c, hm, locale)

				kc1, _ := core.NewCardFromString("A♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(kc1.In((*kcc))).Should(Equal(true))
				Expect(len((*kcc))).Should(Equal(1))
			})
		})
		When("main card is 6♦", func() {
			c, _ := core.NewCardFromString("6♦", locale)
			kcc, err := core.ComputeKarmaCards(c, hm, locale)

			kc1, _ := core.NewCardFromString("9♣", locale)
			kc2, _ := core.NewCardFromString("3♠", locale)

			It("should return array with 2 cards", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(kc1.In((*kcc))).Should(Equal(true))
				Expect(kc2.In((*kcc))).Should(Equal(true))
				Expect(len((*kcc))).Should(Equal(2))
			})
		})
	})

	Describe("ComputePlanetCycles", func() { // TODO
	})

	Describe("ComputeCalendar", func() {
		When("brithday is 2000-09-05 and request for 2019 year", func() {
			planets := locale.GetPlanets()

			b := time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
			cal, err := core.ComputeCalendar(b, od, planets, 2019, mm)

			It("should not raise an error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return arrays of weeks with 2 elements", func() {
				Expect(len((cal.Weeks))).Should(Equal(2))
			})
			It("should return first week with valid set of cards", func() {
				hrw := []string{"9♦", "9♠", "8♠", "A♥", "8♦", "Q♦", "4♦"}
				vrw := []string{"K♥", "6♥", "3♥", "2♦", "10♠", "4♦"}
				dsw := time.Date(2017, time.December, 5, 0, 0, 0, 0, time.UTC)
				dew := time.Date(2019, time.August, 26, 0, 0, 0, 0, time.UTC)

				for i := 0; i < 7; i++ {
					hcw, _ := core.NewCardFromString(hrw[i], locale)
					vcw, _ := core.NewCardFromString(vrw[i], locale)

					Expect(cal.Weeks[0].Days[i].Cards.H).Should(Equal(hcw))
					Expect(cal.Weeks[0].Days[i].Cards.V).Should(Equal(vcw))
				}
				Expect(cal.Weeks[0].Start).Should(Equal(dsw))
				Expect(cal.Weeks[0].End).Should(Equal(dew))
			})
			It("should return second week with valid set of cards", func() {
				hrw := []string{"J♠", "10♠", "4♥", "9♠", "8♥", "K♣", "Q♦"}
				vrw := []string{"A♠", "J♥", "5♣", "J♦", "2♦", "6♣", "Q♦"}
				dsw := time.Date(2019, time.August, 27, 0, 0, 0, 0, time.UTC)
				dew := time.Date(2021, time.May, 17, 0, 0, 0, 0, time.UTC)

				for i := 0; i < 7; i++ {
					hcw, _ := core.NewCardFromString(hrw[i], locale)
					vcw, _ := core.NewCardFromString(vrw[i], locale)

					Expect(cal.Weeks[0].Days[i].Cards.H).Should(Equal(hcw))
					Expect(cal.Weeks[0].Days[i].Cards.V).Should(Equal(vcw))
				}
				Expect(cal.Weeks[0].Start).Should(Equal(dsw))
				Expect(cal.Weeks[0].End).Should(Equal(dew))
			})
		})
	})
})
