package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Deck", func() {
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

		locale = "ru-RU"

		od = core.NewOrderedDeck(locale)
		om = core.NewOriginMatrix(&origin, od)

		mm = core.NewBunchOfYearMatrices(om, od)
	)

	// age: 40, 3♦ -> 9♣
	// age: 37, J♣ -> K♠
	// age: 40, 9♣ -> J♦
	Describe("resolve longterm card", func() {
		Context("when card is 3♦ and age is 40", func() {
			It("should return 9♣", func() {
				b := time.Date(1986, time.April, 15, 0, 0, 0, 0, time.UTC)
				p, _ := core.NewPerson(b, od, mm, om, nil, nil)
				Expect(p.Cards.Longterm).To(Equal(core.NewCardFromNumber(22, locale))
				// Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
