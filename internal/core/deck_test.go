package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Deck", func() {
	var (
		os = [52]uint8{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
			14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
			27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
		}
	)

	Describe("Get card by number", func() {
		Context("With valid index", func() {
			It("should return a valid card", func() {
				var c *core.Card

				for i := uint8(1); i < 52; i++ {
					c, _ = core.NewCardFromNumber(i, locale)
					Expect(od.FindCardByNumber(i)).To(Equal(c))
				}
			})
		})

		Context("With invalid index", func() {
			It("should return an error", func() {
				for _, v := range []int{-1, 0, 53} {
					_, err := od.FindCardByNumber(uint8(v))
					Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
				}
			})
		})
	})

	Describe("Get card by birthday", func() {
		Context("With valid time.Time", func() {
			It("should return a valid card", func() {
				b := time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)

				for i := 1; i <= 365; i++ {
					c, err := od.FindCardByBirthday(b.AddDate(0, 0, i))
					Expect(err).ShouldNot(HaveOccurred())
					Expect(c).Should(BeAssignableToTypeOf(&core.Card{}))
				}
			})
		})
		Context("With birthday at XXXX-12-31", func() {
			It("should raise an error", func() {
				b := time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
				_, err := od.FindCardByBirthday(b)
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	// TODO: test indexOf

	Describe("Build deck", func() {
		Context("from valid slice", func() {
			It("should return a deck", func() {
				Expect(core.NewDeckFromSlice(os, od)).To(Equal(od))
			})
		})

		Context("from invalid slice", func() {
			It("should raise an error", func() {
				s := os
				s[0] = 255
				_, err := core.NewDeckFromSlice(s, od)
				Expect(err).Should(HaveOccurred())
				// Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
			})
		})
	})
})
