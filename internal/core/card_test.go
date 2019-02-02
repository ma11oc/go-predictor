package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Card", func() {
	var c *core.Card

	locales := core.MustLoadLocales("../../locales/ru-RU.yaml")
	lang := language.Make("ru-RU")

	Describe("Get card", func() {
		Context("With number 1", func() {
			It("should return A of hearts", func() {
				c, _ = core.NewCardFromNumber(1, lang, locales)
				Expect(c.ID).To(Equal(uint8(1)))
				Expect(c.Rank).To(Equal("A"))
				Expect(c.Suit).To(Equal("\u2665"))
			})
		})

		Context("With number 52", func() {
			It("should return K of spides", func() {
				c, _ = core.NewCardFromNumber(52, lang, locales)
				Expect(c.ID).To(Equal(uint8(52)))
				Expect(c.Rank).To(Equal("K"))
				Expect(c.Suit).To(Equal("\u2660"))
			})
		})

		Context("With valid number", func() {
			It("should return a valid card", func() {
				for i := uint8(1); i < 53; i++ {
					c, _ = core.NewCardFromNumber(i, lang, locales)
					Expect(c).Should(BeAssignableToTypeOf(&core.Card{}))
				}
			})
		})

		Context("With invalid number", func() {
			It("should return an error", func() {
				for _, v := range []int{-1, 0, 53} {
					_, err := core.NewCardFromNumber(uint8(v), lang, locales)
					Expect(err).Should(HaveOccurred(), "current card number is: %v", v)
				}
			})
		})
	})
})
