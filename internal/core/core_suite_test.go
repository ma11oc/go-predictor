package core_test

import (
	"testing"

	"bitbucket.org/shchukin_a/go-predictor/internal/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"
)

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

	locales = core.BuildLocales("../../locales/ru-RU.yaml")
	lang    = language.Make("ru-RU")
	locale  = locales[lang]

	od = locale.GetOrderedDeck()
	om = locale.GetOriginMatrix()
	hm = locale.GetHumansMatrix()
	am = locale.GetAngelsMatrix()
	mm = locale.GetYearMatrices()
)

func TestCore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Core Suite")
}
