package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Deck", func() {
	var (
		od = core.NewOrderedDeck()
		os = [52]uint8{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
			14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
			27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
		}
	)

	Describe("Get card", func() {
		Context("With valid index", func() {
			It("should return a valid card", func() {
				var c *core.Card

				for i := uint8(1); i < 52; i++ {
					c, _ = core.NewCardFromNumber(i)
					Expect(od.GetCardByNumber(i)).To(Equal(c))
				}
			})
		})

		Context("With invalid index", func() {
			It("should return an error", func() {
				for _, v := range []int{-1, 0, 53} {
					_, err := od.GetCardByNumber(uint8(v))
					Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
				}
			})
		})
	})

	// TODO: test indexOf

	Describe("Build deck", func() {
		Context("from valid slice", func() {
			It("should return a deck", func() {
				Expect(core.NewDeckFromSlice(os)).To(Equal(od))
			})
		})

		Context("from invalid slice", func() {
			It("should raise an error", func() {
				s := os
				s[0] = 255
				_, err := core.NewDeckFromSlice(s)
				Expect(err).Should(HaveOccurred())
				// Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
			})
		})
	})
})
