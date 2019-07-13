package core_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
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

				Expect(err).ShouldNot(HaveOccurred())
				Expect(core.NewCardFromString("K♠", locale)).To(Equal(p.BaseCards["main"]))
				Expect(core.NewCardFromString("K♠", locale)).To(Equal(p.BaseCards["drain"]))
				Expect(core.NewCardFromString("K♠", locale)).To(Equal(p.BaseCards["source"]))
			})
		})

		Context("when birthday is 5 of September", func() {
			It("should return (Main, Drain, Source) == (6♦, 9♣, 3♠)", func() {
				pp.Birthday = time.Date(2000, time.September, 5, 0, 0, 0, 0, time.UTC)
				p, err := core.NewPerson(pp, locale)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(core.NewCardFromString("6♦", locale)).To(Equal(p.BaseCards["main"]))
				Expect(core.NewCardFromString("9♣", locale)).To(Equal(p.BaseCards["drain"]))
				Expect(core.NewCardFromString("3♠", locale)).To(Equal(p.BaseCards["source"]))
			})
		})
	})

})
