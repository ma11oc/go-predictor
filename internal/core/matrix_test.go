package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Matrix", func() {
	var (
		origin = [52]uint8{
			3, 14, 25, 49, 18, 29, 40,
			7, 33, 44, 11, 22, 48, 2,
			13, 39, 6, 17, 28, 50, 21,
			32, 43, 10, 36, 47, 1, 27,
			38, 5, 16, 42, 9, 20, 31,
			51, 24, 35, 46, 15, 26, 37,
			4, 30, 41, 8, 19, 45, 12,
			23, 34, 52,
		}

		locales = core.MustLoadLocales("../../locales/ru-RU.yaml")
		lang    = language.Make("ru-RU")

		od = core.NewOrderedDeck(lang, locales)
		om = core.NewOriginMatrix(&origin, od)
	)

	BeforeSuite(func() {
		// am := core.NewAngelsMatrix(om, od)
		// ym := core.NewZeroYearMatrix(om)
	})

	BeforeEach(func() {
	})

	Describe("Build origin matrix", func() {
		Context("with valid origin matrix and ordered deck", func() {
			It("should return a valid matrix", func() {
				md, _ := core.NewDeckFromSlice(origin, od)
				Expect(om.Decks.Main).To(Equal(md))
				Expect(om.Decks.Drain).To(Equal(od))
				Expect(om.Decks.Source).To(Equal(md))
			})
		})
		// TODO: raise an error
	})

	Describe("Build human matrix", func() {
		Context("with valid origin matrix and ordered deck", func() {
			It("should return a valid matrix", func() {
				hm := core.NewHumansMatrix(om, od)
				sd, _ := core.NewDeckFromSlice([52]uint8{
					25, 2, 36, 12, 17, 38, 15, 40, 9, 30, 11, 32, 45,
					14, 48, 46, 29, 6, 27, 23, 21, 42, 4, 44, 51, 19,
					3, 1, 35, 18, 39, 37, 33, 50, 16, 34, 10, 31, 8,
					13, 47, 24, 49, 5, 26, 7, 28, 41, 22, 43, 20, 52,
				}, od)
				Expect(hm.Decks.Main).To(Equal(om.Decks.Main))
				Expect(hm.Decks.Drain).To(Equal(od))
				Expect(hm.Decks.Source).To(Equal(sd))
			})
		})
		// TODO: raise an error

	})

	Describe("Build angels matrix", func() {
		Context("with valid origin matrix and ordered deck", func() {
			It("should return a valid matrix", func() {
				am := core.NewAngelsMatrix(om, od)
				dd, _ := core.NewDeckFromSlice([52]uint8{
					27, 14, 1, 43, 30, 17, 8, 46, 33, 24, 11, 49, 15,
					2, 40, 31, 18, 5, 47, 34, 21, 12, 50, 37, 3, 41,
					28, 19, 6, 44, 35, 22, 9, 51, 38, 25, 42, 29, 16,
					7, 45, 32, 23, 10, 48, 39, 26, 13, 4, 20, 36, 52,
				}, od)
				Expect(am.Decks.Main).To(Equal(od))
				Expect(am.Decks.Drain).To(Equal(dd))
				Expect(am.Decks.Source).To(Equal(om.Decks.Main))
			})
		})
		// TODO: raise an error

	})

	Describe("Build next year matrix", func() {
		Context("with valid origin matrix", func() {
			It("should return a valid matrix", func() {
				zy := core.NewZeroYearMatrix(om)
				Expect(zy.Decks.Main).To(Equal(om.Decks.Main))
				Expect(zy.Decks.Drain).To(Equal(om.Decks.Drain))
				Expect(zy.Decks.Source).To(Equal(om.Decks.Source))
			})
		})

		When("year belongs to the interval [1, 89]", func() {
			It("should return a valid matrix", func() {
				current := core.NewZeroYearMatrix(om)
				for i := uint8(1); i <= 89; i++ {
					next, _ := current.Next(om, od)
					current = next
					// Expect(next.Decks.Main).To(Equal(om.Decks.Main))
					Expect(next.Decks.Drain).To(Equal(od))
					Expect(next.Decks.Source).To(Equal(om.Decks.Main))
				}
			})
		})
		// TODO: raise an error

	})
})
