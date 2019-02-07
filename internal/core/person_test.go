package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Deck", func() {
	var (
		pc = &core.PersonConfig{}
	)

	Describe("resolve longterm card", func() {
		Context("when card is 3♦ and age is 40", func() {
			It("should return 9♣", func() {
				pc.Birthday = time.Date(time.Now().Year()-41, time.March, 20, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(22, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is J♣ and age is 37", func() {
			It("should return K♠", func() {
				pc.Birthday = time.Date(time.Now().Year()-38, time.March, 25, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(52, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 9♣ and age is 40", func() {
			It("should return J♦", func() {
				pc.Birthday = time.Date(time.Now().Year()-41, time.March, 27, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(37, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 6♦ and age is 66", func() {
			It("should return 10♠", func() {
				pc.Birthday = time.Date(time.Now().Year()-66, time.January, 21, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(49, locale)).To(Equal(p.Cards.Longterm))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("resolve pluto/result card", func() {
		Context("when card is Q♦ and age is 31", func() {
			It("should return pluto as 2♦ and pluto result as 2♠", func() {
				pc.Birthday = time.Date(time.Now().Year()-31, time.January, 15, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(28, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(41, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is K♥ and age is 31", func() {
			It("should return pluto as 4♦ and pluto result as 5♠", func() {
				pc.Birthday = time.Date(time.Now().Year()-32, time.September, 24, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(30, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(44, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 10♥ and age is 31", func() {
			It("should return pluto as 2♥ and pluto result as 7♥", func() {
				pc.Birthday = time.Date(time.Now().Year()-32, time.November, 23, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(2, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(7, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when card is 9♦ and age is 67", func() {
			It("should return pluto as 2♥ and pluto result as 7♥", func() {
				pc.Birthday = time.Date(time.Now().Year()-68, time.September, 2, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pc, locale)
				Expect(core.NewCardFromNumber(49, locale)).To(Equal(p.Cards.Pluto))
				Expect(core.NewCardFromNumber(4, locale)).To(Equal(p.Cards.PlutoResult))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
