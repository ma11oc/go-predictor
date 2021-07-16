package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ma11oc/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Card", func() {
	var c *core.Card

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
