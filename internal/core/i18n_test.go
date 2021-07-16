package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ma11oc/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Card", func() {
	var c *core.Card

	Describe("NewCardFromNumber", func() {
		Context("with number 1", func() {
			It("should return A♥", func() {
				c, _ = core.NewCardFromNumber(1, locale)
				Expect(c.ID).To(Equal(uint8(1)))
				Expect(c.Rank).To(Equal("A"))
				Expect(c.Suit).To(Equal("\u2665"))
			})
		})

		Context("with number 52", func() {
			It("should return K♠", func() {
				c, _ = core.NewCardFromNumber(52, locale)
				Expect(c.ID).To(Equal(uint8(52)))
				Expect(c.Rank).To(Equal("K"))
				Expect(c.Suit).To(Equal("\u2660"))
			})
		})

		Context("with valid number", func() {
			It("should return a valid card", func() {
				for i := uint8(1); i < 53; i++ {
					c, _ = core.NewCardFromNumber(i, locale)
					Expect(c).Should(BeAssignableToTypeOf(&core.Card{}))
				}
			})
		})

		Context("with invalid number", func() {
			It("should return an error", func() {
				for _, v := range []int{-1, 0, 53} {
					_, err := core.NewCardFromNumber(uint8(v), locale)
					Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
				}
			})
		})
	})

	Describe("NewCardFromString", func() {
		Context("with string A♥", func() {
			It("should return A♥", func() {
				c, _ = core.NewCardFromString("A♥", locale)
				Expect(c.ID).To(Equal(uint8(1)))
				Expect(c.Rank).To(Equal("A"))
				Expect(c.Suit).To(Equal("\u2665"))
			})
		})

		Context("with valid string", func() {
			It("should return a valid card", func() {
				for _, rank := range []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
					for _, suit := range []string{"♥", "♣", "♦", "♠"} {
						c, _ = core.NewCardFromString(rank+suit, locale)
						Expect(c).Should(BeAssignableToTypeOf(&core.Card{}))
					}
				}
			})
		})
		// TODO: check special case with Joker
	})
})
