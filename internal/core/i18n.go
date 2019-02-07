package core

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/go-validator/validator"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

func init() {
	validator.SetValidationFunc("langstr", isParsableLanguageTag)
}

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
)

type Locale struct {
	Lang string `yaml:"lang" validate:"nonzero,langstr"`
	Base struct {
		od *Deck
		om *Matrix
		hm *Matrix
		am *Matrix
		mm *[90]*YearMatrix
		pc *[7][54]*PlanetCycle
	}
	Matrices []struct {
		ID string `yaml:"id"`
	}
	Exceptions struct {
		Joker *Card
	}
	Cards []Card `yaml:"cards" validate:"len=52"`
}

func (l Locale) GetOrderedDeck() *Deck {
	return l.Base.od
}

func (l Locale) GetYearMatrices() *[90]*YearMatrix {
	return l.Base.mm
}

func (l Locale) GetOriginMatrix() *Matrix {
	return l.Base.om
}

func (l Locale) GetHumansMatrix() *Matrix {
	return l.Base.hm
}

func (l Locale) GetAngelsMatrix() *Matrix {
	return l.Base.am
}

type Locales map[language.Tag]*Locale

// Make sure a value of a field is parsable by language.Parse()
func isParsableLanguageTag(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return validator.ErrUnsupported

	}
	if _, err := language.Parse(st.String()); err != nil {
		return err
	}

	return nil
}

// loadtranslation
//
/*
 * type Locale struct {
 *     Lang     string `yaml:"lang" validate:"nonzero,langstr"`
 *     Matrices []struct {
 *         ID string `yaml:"id"`
 *     }
 *     Exceptions struct {
 *         Joker *Card
 *     }
 *     Cards []Card `yaml:"cards" validate:"len=52"`
 * }
 */

func (l *Locale) GetCardByID(id uint8) (*Card, error) {
	if id <= 0 || id > 52 {
		return nil, fmt.Errorf("Wrong id has been specified: %v", id)
	}

	card := l.Cards[id-1]

	if card.ID != id {
		return nil, fmt.Errorf("Locale seems broken: index != id-1 (index: %v, id: %v)", id-1, card.ID)
	}

	return &card, nil
}

func (l *Locale) Validate() error {
	if errs := validator.Validate(l); errs != nil {
		return errs
	}
	return nil
}

func NewLocale(p string) (*Locale, error) {
	var content []byte
	var err error

	if content, err = ioutil.ReadFile(p); err != nil {
		return nil, fmt.Errorf("unable to read locale file: %v", err)
	}

	loc := &Locale{}

	if err = yaml.Unmarshal([]byte(content), &loc); err != nil {
		return nil, fmt.Errorf("unable to unmarshall locale: %v", err)
	}

	if err = loc.Validate(); err != nil {
		return nil, fmt.Errorf("locale %v is invalid: %v", p, err)
	}

	loc.Base.od = NewOrderedDeck(loc)
	loc.Base.om = NewOriginMatrix(&origin, loc.Base.od)
	loc.Base.hm = NewHumansMatrix(loc.Base.om, loc.Base.od)
	loc.Base.am = NewAngelsMatrix(loc.Base.om, loc.Base.od)
	loc.Base.mm = NewBunchOfYearMatrices(loc.Base.om, loc.Base.od)
	loc.Base.pc = NewBunchOfPlanetCycles()

	return loc, nil
}

func BuildLocales(paths ...string) Locales {
	var loc *Locale
	var err error

	ll := make(map[language.Tag]*Locale, len(paths))

	for _, p := range paths {
		if loc, err = NewLocale(p); err != nil {
			panic(err)
		}

		// Lang has been validated before, so there is not need to parse it
		lang := language.Make(loc.Lang)

		ll[lang] = loc
	}

	return ll
}
