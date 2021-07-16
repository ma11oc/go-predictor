package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ma11oc/go-predictor/internal/core"
)

var _ = Describe("Internal/Core/Person", func() {
	var (
		pp = &core.PersonProfile{
			Name:     "John",
			Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			Gender:   core.Male,
			Features: core.Feature(0x0),
		}
	)

	Describe("NewPerson", func() {
		Context("when birthday is 1 of January", func() {
			It("should return (Main, Drain, Source) == (K♠, K♠, K♠)", func() {
				p, err := core.NewPerson(pp, locale)

				cw, _ := core.NewCardFromString("K♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.BaseCards["main"]).To(Equal(cw))
				Expect(p.BaseCards["drain"]).To(Equal(cw))
				Expect(p.BaseCards["source"]).To(Equal(cw))
			})
		})

		Context("when birthday is 5 of September", func() {
			It("should return (Main, Drain, Source) == (6♦, 9♣, 3♠)", func() {
				pp.Birthday = time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pp, locale)

				mcw, _ := core.NewCardFromString("6♦", locale)
				dcw, _ := core.NewCardFromString("9♣", locale)
				scw, _ := core.NewCardFromString("3♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.BaseCards["main"]).To(Equal(mcw))
				Expect(p.BaseCards["drain"]).To(Equal(dcw))
				Expect(p.BaseCards["source"]).To(Equal(scw))
			})
		})

		Context("when birthday is 5 of September", func() {
			It("should return return a valid person, when age is 33", func() {
				pp.Birthday = time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				pp.Age = 33
				p, err := core.NewPerson(pp, locale)

				mcw, _ := core.NewCardFromString("6♦", locale)
				dcw, _ := core.NewCardFromString("9♣", locale)
				scw, _ := core.NewCardFromString("3♠", locale)
				lcw, _ := core.NewCardFromString("J♦", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.BaseCards["main"]).To(Equal(mcw))
				Expect(p.BaseCards["drain"]).To(Equal(dcw))
				Expect(p.BaseCards["source"]).To(Equal(scw))
				Expect(p.BaseCards["longterm"]).To(Equal(lcw))

				hrw := []string{"4♠", "J♣", "10♦", "7♦", "2♠", "J♦", "9♠"}
				vrw := []string{"Q♣", "6♥", "5♥", "A♠", "2♦", "9♠"}
				// TODO: check v[3]

				for i := range core.PlanetsOrder {
					vc := &core.Card{}
					hc := p.PlanetCycles[i].Cards.H
					vcw := &core.Card{}

					if p.PlanetCycles[i].Cards.V != nil {
						vc = p.PlanetCycles[i].Cards.V
						vcw, _ = core.NewCardFromString(vrw[i], locale)

						Expect(vc).To(Equal(vcw))
					} else {
						vc = nil
						Expect(vc).To(BeNil())
					}

					hcw, _ := core.NewCardFromString(hrw[i], locale)

					Expect(hc).To(Equal(hcw))
				}

			})
			It("should return return a valid person, when age is 34", func() {
				pp.Birthday = time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				pp.Age = 34
				p, err := core.NewPerson(pp, locale)

				mcw, _ := core.NewCardFromString("6♦", locale)
				dcw, _ := core.NewCardFromString("9♣", locale)
				scw, _ := core.NewCardFromString("3♠", locale)
				lcw, _ := core.NewCardFromString("Q♥", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.BaseCards["main"]).To(Equal(mcw))
				Expect(p.BaseCards["drain"]).To(Equal(dcw))
				Expect(p.BaseCards["source"]).To(Equal(scw))
				Expect(p.BaseCards["longterm"]).To(Equal(lcw))

				hrw := []string{"4♣", "5♣", "7♦", "10♣", "J♥", "Q♣", "Q♦"}
				vrw := []string{"9♦", "3♣", "K♣", "A♠", "8♠", "Q♦"}

				for i := range core.PlanetsOrder {
					vc := &core.Card{}
					hc := p.PlanetCycles[i].Cards.H
					vcw := &core.Card{}

					if p.PlanetCycles[i].Cards.V != nil {
						vc = p.PlanetCycles[i].Cards.V
						vcw, _ = core.NewCardFromString(vrw[i], locale)

						Expect(vc).To(Equal(vcw))
					} else {
						vc = nil
						Expect(vc).To(BeNil())
					}

					hcw, _ := core.NewCardFromString(hrw[i], locale)

					Expect(hc).To(Equal(hcw))
				}
			})
			It("should return return a valid person, when age is 35", func() {
				pp.Birthday = time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				pp.Age = 35
				p, err := core.NewPerson(pp, locale)

				mcw, _ := core.NewCardFromString("6♦", locale)
				dcw, _ := core.NewCardFromString("9♣", locale)
				scw, _ := core.NewCardFromString("3♠", locale)
				lcw, _ := core.NewCardFromString("4♣", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.BaseCards["main"]).To(Equal(mcw))
				Expect(p.BaseCards["drain"]).To(Equal(dcw))
				Expect(p.BaseCards["source"]).To(Equal(scw))
				Expect(p.BaseCards["longterm"]).To(Equal(lcw))

				hrw := []string{"9♣", "7♥", "6♣", "8♣", "Q♠", "A♦", "10♣"}
				vrw := []string{"4♠", "9♠", "2♦", "3♣", "J♣", "5♠", "10♣"}

				for i := range core.PlanetsOrder {
					vc := &core.Card{}
					hc := p.PlanetCycles[i].Cards.H
					vcw := &core.Card{}

					if p.PlanetCycles[i].Cards.V != nil {
						vc = p.PlanetCycles[i].Cards.V
						vcw, _ = core.NewCardFromString(vrw[i], locale)

						Expect(vc).To(Equal(vcw))
					} else {
						vc = nil
						Expect(vc).To(BeNil())
					}

					hcw, _ := core.NewCardFromString(hrw[i], locale)

					Expect(hc).To(Equal(hcw))
				}
			})
			It("should return return a valid person, when age is 66", func() {
				pp.Birthday = time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				pp.Age = 66
				p, err := core.NewPerson(pp, locale)

				mcw, _ := core.NewCardFromString("6♦", locale)
				dcw, _ := core.NewCardFromString("9♣", locale)
				scw, _ := core.NewCardFromString("3♠", locale)
				lcw, _ := core.NewCardFromString("10♠", locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.BaseCards["main"]).To(Equal(mcw))
				Expect(p.BaseCards["drain"]).To(Equal(dcw))
				Expect(p.BaseCards["source"]).To(Equal(scw))
				Expect(p.BaseCards["longterm"]).To(Equal(lcw))

				hrw := []string{"7♥", "3♣", "Q♦", "K♠", "5♣", "A♣", "4♣"}
				vrw := []string{"J♦", "9♦", "5♥", "5♦", "9♣", "3♥"}

				for i := range core.PlanetsOrder {
					vc := &core.Card{}
					hc := p.PlanetCycles[i].Cards.H
					vcw := &core.Card{}

					if p.PlanetCycles[i].Cards.V != nil {
						vc = p.PlanetCycles[i].Cards.V
						vcw, _ = core.NewCardFromString(vrw[i], locale)

						Expect(vc).To(Equal(vcw))
					} else {
						vc = nil
						Expect(vc).To(BeNil())
					}

					hcw, _ := core.NewCardFromString(hrw[i], locale)

					Expect(hc).To(Equal(hcw))
				}
			})
		})
	})

})
