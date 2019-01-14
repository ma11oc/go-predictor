package core_test

import (
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

		od = core.NewOrderedDeck("ru-RU")
		om = core.NewOriginMatrix(&origin, od)

		mm = core.NewBunchOfYearMatrices(om, od)

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
					c, _ = core.NewCardFromNumber(i, "ru-RU")
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

	Describe("GetHRow", func() {
		Context("for 3♦ (29)", func() {
			It("should return a valid slice of cards", func() {
				/* GetHRow for 3♦ (29)
				* 0: 40 | A♠
				* 1:  7 | 7♥
				* 2: 33 | 7♦
				* 3: 44 | 5♠
				* 4: 11 | J♥
				* 5: 22 | 9♣
				* 6: 48 | 9♠
				 */

				card, _ := core.NewCardFromNumber(29, "ru-RU")
				row, err := mm[0].Decks.Main.GetHRow(card)

				s := [7]uint8{40, 7, 33, 44, 11, 22, 48}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromNumber(v, "ru-RU")
					Expect(row[i]).To(Equal(c))
				}
			})
		})

		Context("for 6♠ (45)", func() {
			It("should return a valid slice of cards", func() {
				/* GetHRow for 6♠ (45)
				* 0: 12 | Q♥
				* 1: 23 | 10♣
				* 2: 34 | 8♦
				* 3: 52 | K♠
				* 4:  3 | 3♥
				* 5: 14 | A♣
				* 6: 25 | Q♣
				 */

				card, _ := core.NewCardFromNumber(45, "ru-RU")
				row, err := mm[0].Decks.Main.GetHRow(card)

				s := [7]uint8{12, 23, 34, 52, 3, 14, 25}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromNumber(v, "ru-RU")
					Expect(row[i]).To(Equal(c))
				}
			})
		})
		Context("for 8♦ (34)", func() {
			It("should return a valid slice of cards", func() {
				/* GetHRow for 8♦ (34)
				* 0: 52 | K♠
				* 1:  3 | 3♥
				* 2: 14 | A♣
				* 3: 25 | Q♣
				* 4: 49 | 10♠
				* 5: 18 | 5♣
				* 6: 29 | 3♦
				 */

				card, _ := core.NewCardFromNumber(34, "ru-RU")
				row, err := mm[0].Decks.Main.GetHRow(card)

				s := [7]uint8{52, 3, 14, 25, 49, 18, 29}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromNumber(v, "ru-RU")
					Expect(row[i]).To(Equal(c))
				}
			})
		})
		Context("for 4♣ (17)", func() {
			It("should return a valid slice of cards", func() {
				/* GetHRow for 4♣ (17)
				* 0: 28 | 2♦
				* 1: 50 | J♠
				* 2: 21 | 8♣
				* 3: 32 | 6♦
				* 4: 43 | 4♠
				* 5: 10 | 10♥
				* 6: 36 | 10♦
				 */

				card, _ := core.NewCardFromNumber(17, "ru-RU")
				row, err := mm[0].Decks.Main.GetHRow(card)

				s := [7]uint8{28, 50, 21, 32, 43, 10, 36}

				Expect(err).ShouldNot(HaveOccurred())

				for i, v := range s {
					c, _ := core.NewCardFromNumber(v, "ru-RU")
					Expect(row[i]).To(Equal(c))
				}
			})
		})
	})

	Describe("GetVRow", func() {
		Context("for 3♦ (29)", func() {
			It("should return a valid slice of cards", func() {
				/* GetVRow for 3♦ (29)
				* 0: 45 | 6♠
				* 1: 26 | K♣
				* 2: 20 | 7♣
				* 3:  1 | A♥
				* 4: 50 | J♠
				* 5: 48 | 9♠
				 */

				card, _ := core.NewCardFromNumber(29, "ru-RU")
				row, err := mm[0].Decks.Main.GetVRow(card)

				s := [7]uint8{45, 26, 20, 1, 50, 48}

				Expect(err).ShouldNot(HaveOccurred())
				// Expect(row[6]).To(Equal(nil))

				for i, v := range s {
					c, _ := core.NewCardFromNumber(v, "ru-RU")
					Expect(row[i]).To(Equal(c))
				}
			})
		})
		Context("for 4♣ (17)", func() {
			It("should return a valid slice of cards", func() {
				/* GetVRow for 4♣ (17)
				* 0: 11 | J♥
				* 1: 49 | 10♠
				* 2: 34 | 8♦
				* 3: 8  | 8♥
				* 4: 46 | 7♠
				* 5: 42 | 3♠
				* 6: 36 | 10♦
				 */

				card, _ := core.NewCardFromNumber(17, "ru-RU")
				row, err := mm[0].Decks.Main.GetVRow(card)

				s := [7]uint8{11, 49, 34, 8, 46, 42, 36}

				Expect(err).ShouldNot(HaveOccurred())
				// Expect(row[6]).To(Equal(nil))

				for i, v := range s {
					c, _ := core.NewCardFromNumber(v, "ru-RU")
					Expect(row[i]).To(Equal(c))
				}
			})
		})
		Context("for all the cards in all the years", func() {
			It("should not raise an error", func() {
				for i := uint8(0); i < 90; i++ {
					for j := uint8(1); j <= 52; j++ {
						c, _ := core.NewCardFromNumber(j, "ru-RU")
						row, err := mm[i].Decks.Main.GetVRow(c)

						Expect(row).Should(BeAssignableToTypeOf([7]*core.Card{}))
						Expect(err).ShouldNot(HaveOccurred())
					}
				}
			})
		})
	})

	/*
	 * for i := uint8(0); i < 90; i++ {
	 *     for j := uint8(1); j <= 52; j++ {
	 *         c, _ := core.NewCardFromNumber(j)
	 *         mm[i].Decks.Main.GetVRow(c)
	 *     }
	 * }
	 */

})
