package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Card", func() {
	var c *core.Card

	Describe("New card", func() {
		Context("With number 1", func() {
			It("should return A of hearts", func() {
				c, _ = core.NewCardFromNumber(1, locale)
				Expect(c.ID).To(Equal(uint8(1)))
				Expect(c.Rank).To(Equal("A"))
				Expect(c.Suit).To(Equal("\u2665"))
			})
		})

		Context("With number 52", func() {
			It("should return K of spides", func() {
				c, _ = core.NewCardFromNumber(52, locale)
				Expect(c.ID).To(Equal(uint8(52)))
				Expect(c.Rank).To(Equal("K"))
				Expect(c.Suit).To(Equal("\u2660"))
			})
		})

		Context("With valid number", func() {
			It("should return a valid card", func() {
				for i := uint8(1); i < 53; i++ {
					c, _ = core.NewCardFromNumber(i, locale)
					Expect(c).Should(BeAssignableToTypeOf(&core.Card{}))
				}
			})
		})

		Context("With invalid number", func() {
			It("should return an error", func() {
				for _, v := range []int{-1, 0, 53} {
					_, err := core.NewCardFromNumber(uint8(v), locale)
					Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
				}
			})
		})
	})

	Describe("GetBirthdays", func() {
		Context("for K♠", func() {
			It("should return the array with 1 element", func() {
				c, _ = core.NewCardFromNumber(52, locale)
				bdays, err := c.GetBirthdays()

				want := []time.Time{
					time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				}

				Expect(want).To(Equal(bdays))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("for 10♥", func() {
			It("should return the array with 6 elements", func() {
				c, _ = core.NewCardFromNumber(10, locale)
				bdays, err := c.GetBirthdays()

				want := []time.Time{
					time.Date(2000, time.July, 31, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.August, 29, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.September, 27, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.October, 25, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.November, 23, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.December, 21, 0, 0, 0, 0, time.UTC),
				}

				Expect(want).To(Equal(bdays))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("for 3♥", func() {
			It("should return the array with 2 elements", func() {
				c, _ = core.NewCardFromNumber(3, locale)
				bdays, err := c.GetBirthdays()

				want := []time.Time{
					time.Date(2000, time.November, 30, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.December, 28, 0, 0, 0, 0, time.UTC),
				}

				Expect(want).To(Equal(bdays))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("for 10♣", func() {
			It("should return the array with 12 elements", func() {
				c, _ = core.NewCardFromNumber(23, locale)
				bdays, err := c.GetBirthdays()

				want := []time.Time{
					time.Date(2000, time.January, 30, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.February, 28, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.March, 26, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.April, 24, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.May, 22, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.June, 20, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.July, 18, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.August, 16, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.September, 14, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.October, 12, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2000, time.December, 8, 0, 0, 0, 0, time.UTC),
				}

				Expect(want).To(Equal(bdays))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
