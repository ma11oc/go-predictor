package core_test

/*
 * hm | main[52]: K♠  | drain[52]: K♠  | source[52]: K♠
 * hm | main[34]: 8♦  | drain[51]: Q♠  | source[20]: 7♣
 * hm | main[23]: 10♣ | drain[50]: J♠  | source[43]: 4♠
 *
 * hm | main[40]: A♠  | drain[ 7]: 7♥  | source[15]: 2♣
 * hm | main[29]: 3♦  | drain[ 6]: 6♥  | source[38]: Q♦
 * hm | main[18]: 5♣  | drain[ 5]: 5♥  | source[17]: 4♣
 * hm | main[49]: 10♠ | drain[ 4]: 4♥  | source[12]: Q♥
 * hm | main[25]: Q♣  | drain[ 3]: 3♥  | source[36]: 10♦
 * hm | main[14]: A♣  | drain[ 2]: 2♥  | source[ 2]: 2♥
 * hm | main[ 3]: 3♥  | drain[ 1]: A♥  | source[25]: Q♣
 *
 * hm | main[ 2]: 2♥  | drain[14]: A♣  | source[14]: A♣
 * hm | main[48]: 9♠  | drain[13]: K♥  | source[45]: 6♠
 * hm | main[22]: 9♣  | drain[12]: Q♥  | source[32]: 6♦
 * hm | main[11]: J♥  | drain[11]: J♥  | source[11]: J♥
 * hm | main[44]: 5♠  | drain[10]: 10♥ | source[30]: 4♦
 * hm | main[33]: 7♦  | drain[ 9]: 9♥  | source[ 9]: 9♥
 * hm | main[ 7]: 7♥  | drain[ 8]: 8♥  | source[40]: A♠
 * hm | main[21]: 8♣  | drain[21]: 8♣  | source[21]: 8♣
 * hm | main[50]: J♠  | drain[20]: 7♣  | source[23]: 10♣
 * hm | main[28]: 2♦  | drain[19]: 6♣  | source[27]: A♦
 * hm | main[17]: 4♣  | drain[18]: 5♣  | source[ 6]: 6♥
 * hm | main[ 6]: 6♥  | drain[17]: 4♣  | source[29]: 3♦
 * hm | main[39]: K♦  | drain[16]: 3♣  | source[46]: 7♠
 * hm | main[13]: K♥  | drain[15]: 2♣  | source[48]: 9♠
 * hm | main[27]: A♦  | drain[28]: 2♦  | source[ 1]: A♥
 * hm | main[ 1]: A♥  | drain[27]: A♦  | source[ 3]: 3♥
 * hm | main[47]: 8♠  | drain[26]: K♣  | source[19]: 6♣
 * hm | main[36]: 10♦ | drain[25]: Q♣  | source[51]: Q♠
 * hm | main[10]: 10♥ | drain[24]: J♣  | source[44]: 5♠
 * hm | main[43]: 4♠  | drain[23]: 10♣ | source[ 4]: 4♥
 * hm | main[32]: 6♦  | drain[22]: 9♣  | source[42]: 3♠
 * hm | main[31]: 5♦  | drain[35]: 9♦  | source[16]: 3♣
 * hm | main[20]: 7♣  | drain[34]: 8♦  | source[50]: J♠
 * hm | main[ 9]: 9♥  | drain[33]: 7♦  | source[33]: 7♦
 * hm | main[42]: 3♠  | drain[32]: 6♦  | source[37]: J♦
 * hm | main[16]: 3♣  | drain[31]: 5♦  | source[39]: K♦
 * hm | main[ 5]: 5♥  | drain[30]: 4♦  | source[18]: 5♣
 * hm | main[38]: Q♦  | drain[29]: 3♦  | source[35]: 9♦
 * hm | main[37]: J♦  | drain[42]: 3♠  | source[24]: J♣
 * hm | main[26]: K♣  | drain[41]: 2♠  | source[47]: 8♠
 * hm | main[15]: 2♣  | drain[40]: A♠  | source[13]: K♥
 * hm | main[46]: 7♠  | drain[39]: K♦  | source[ 8]: 8♥
 * hm | main[35]: 9♦  | drain[38]: Q♦  | source[31]: 5♦
 * hm | main[24]: J♣  | drain[37]: J♦  | source[10]: 10♥
 * hm | main[51]: Q♠  | drain[36]: 10♦ | source[34]: 8♦
 * hm | main[12]: Q♥  | drain[49]: 10♠ | source[22]: 9♣
 * hm | main[45]: 6♠  | drain[48]: 9♠  | source[41]: 2♠
 * hm | main[19]: 6♣  | drain[47]: 8♠  | source[28]: 2♦
 * hm | main[ 8]: 8♥  | drain[46]: 7♠  | source[ 7]: 7♥
 * hm | main[41]: 2♠  | drain[45]: 6♠  | source[26]: K♣
 * hm | main[30]: 4♦  | drain[44]: 5♠  | source[ 5]: 5♥
 * hm | main[ 4]: 4♥  | drain[43]: 4♠  | source[49]: 10♠
 *
 * am | main[52]: K♠  | drain[52]: K♠  | source[52]: K♠
 * am | main[51]: Q♠  | drain[36]: 10♦ | source[34]: 8♦
 * am | main[50]: J♠  | drain[20]: 7♣  | source[23]: 10♣
 * am | main[ 7]: 7♥  | drain[ 8]: 8♥  | source[40]: A♠
 * am | main[ 6]: 6♥  | drain[17]: 4♣  | source[29]: 3♦
 * am | main[ 5]: 5♥  | drain[30]: 4♦  | source[18]: 5♣
 * am | main[ 4]: 4♥  | drain[43]: 4♠  | source[49]: 10♠
 * am | main[ 3]: 3♥  | drain[ 1]: A♥  | source[25]: Q♣
 * am | main[ 2]: 2♥  | drain[14]: A♣  | source[14]: A♣
 * > am | main[ 1]: A♥  | drain[27]: A♦  | source[ 3]: 3♥
 * am | main[14]: A♣  | drain[ 2]: 2♥  | source[ 2]: 2♥
 * am | main[13]: K♥  | drain[15]: 2♣  | source[48]: 9♠
 * am | main[12]: Q♥  | drain[49]: 10♠ | source[22]: 9♣
 * am | main[11]: J♥  | drain[11]: J♥  | source[11]: J♥
 * am | main[10]: 10♥ | drain[24]: J♣  | source[44]: 5♠
 * am | main[ 9]: 9♥  | drain[33]: 7♦  | source[33]: 7♦
 * am | main[ 8]: 8♥  | drain[46]: 7♠  | source[ 7]: 7♥
 * am | main[21]: 8♣  | drain[21]: 8♣  | source[21]: 8♣
 * am | main[20]: 7♣  | drain[34]: 8♦  | source[50]: J♠
 * am | main[19]: 6♣  | drain[47]: 8♠  | source[28]: 2♦
 * am | main[18]: 5♣  | drain[ 5]: 5♥  | source[17]: 4♣
 * am | main[17]: 4♣  | drain[18]: 5♣  | source[ 6]: 6♥
 * am | main[16]: 3♣  | drain[31]: 5♦  | source[39]: K♦
 * am | main[15]: 2♣  | drain[40]: A♠  | source[13]: K♥
 * am | main[28]: 2♦  | drain[19]: 6♣  | source[27]: A♦
 * am | main[27]: A♦  | drain[28]: 2♦  | source[ 1]: A♥
 * am | main[26]: K♣  | drain[41]: 2♠  | source[47]: 8♠
 * am | main[25]: Q♣  | drain[ 3]: 3♥  | source[36]: 10♦
 * am | main[24]: J♣  | drain[37]: J♦  | source[10]: 10♥
 * am | main[23]: 10♣ | drain[50]: J♠  | source[43]: 4♠
 * am | main[22]: 9♣  | drain[12]: Q♥  | source[32]: 6♦
 * am | main[35]: 9♦  | drain[38]: Q♦  | source[31]: 5♦
 * am | main[34]: 8♦  | drain[51]: Q♠  | source[20]: 7♣
 * am | main[33]: 7♦  | drain[ 9]: 9♥  | source[ 9]: 9♥
 * am | main[32]: 6♦  | drain[22]: 9♣  | source[42]: 3♠
 * am | main[31]: 5♦  | drain[35]: 9♦  | source[16]: 3♣
 * am | main[30]: 4♦  | drain[44]: 5♠  | source[ 5]: 5♥
 * am | main[29]: 3♦  | drain[ 6]: 6♥  | source[38]: Q♦
 * am | main[42]: 3♠  | drain[32]: 6♦  | source[37]: J♦
 * am | main[41]: 2♠  | drain[45]: 6♠  | source[26]: K♣
 * am | main[40]: A♠  | drain[ 7]: 7♥  | source[15]: 2♣
 * am | main[39]: K♦  | drain[16]: 3♣  | source[46]: 7♠
 * am | main[38]: Q♦  | drain[29]: 3♦  | source[35]: 9♦
 * am | main[37]: J♦  | drain[42]: 3♠  | source[24]: J♣
 * am | main[36]: 10♦ | drain[25]: Q♣  | source[51]: Q♠
 * am | main[49]: 10♠ | drain[ 4]: 4♥  | source[12]: Q♥
 * am | main[48]: 9♠  | drain[13]: K♥  | source[45]: 6♠
 * am | main[47]: 8♠  | drain[26]: K♣  | source[19]: 6♣
 * am | main[46]: 7♠  | drain[39]: K♦  | source[ 8]: 8♥
 * am | main[45]: 6♠  | drain[48]: 9♠  | source[41]: 2♠
 * am | main[44]: 5♠  | drain[10]: 10♥ | source[30]: 4♦
 * am | main[43]: 4♠  | drain[23]: 10♣ | source[ 4]: 4♥
 */

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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

		od = core.NewOrderedDeck()
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
				md, _ := core.NewDeckFromSlice(origin)
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
				})
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
				})
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
