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

// TODO: validate planets order
// TODO: validate cards order

// Locale is a fundamental primitive. It contains all the descriptions
// required for any prediction
type Locale struct {
	Lang string `yaml:"lang" validate:"nonzero,langstr"`

	Core struct {
		od *Deck
		om *Matrix
		hm *Matrix
		am *Matrix
		mm *Matrices
		cc *Cycles
	}
	Descriptors struct {
		Cards   map[string]string `yaml:"cards_meta"   validate:"nonzero"`
		Planets map[string]string `yaml:"planets_meta" validate:"nonzero"`
	}

	Exceptions struct {
		Joker *Card
	}

	Cards []Card `yaml:"cards"   validate:"len=52"`

	Planets *Planets
}

// GetOrderedDeck is ordered Deck getter
func (l Locale) GetOrderedDeck() *Deck {
	return l.Core.od
}

// GetYearMatrices is Year Matrices getter
func (l Locale) GetYearMatrices() *Matrices {
	return l.Core.mm
}

// GetOriginMatrix is Origin Matrix getter
func (l Locale) GetOriginMatrix() *Matrix {
	return l.Core.om
}

// GetHumansMatrix is Humans Matrix getter
func (l Locale) GetHumansMatrix() *Matrix {
	return l.Core.hm
}

// GetCycles is Cycles getter
func (l Locale) GetCycles() *Cycles {
	return l.Core.cc
}

// GetPlanets is Planets getter
func (l Locale) GetPlanets() *Planets {
	return l.Planets
}

// GetAngelsMatrix is Angels Matrix getter
func (l Locale) GetAngelsMatrix() *Matrix {
	return l.Core.am
}

// Locales contains all the available and loaded Locales
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

// FindCardByID receives a number (id) of card and returns appropriate *Card from Locale
func (l *Locale) FindCardByID(id uint8) (*Card, error) {
	if id <= 0 || id > 52 {
		return nil, fmt.Errorf("Wrong id has been specified: %v", id)
	}

	card := l.Cards[id-1]

	if card.ID != id {
		return nil, fmt.Errorf("Locale seems broken: index != id-1 (index: %v, id: %v)", id-1, card.ID)
	}

	return &card, nil
}

// FindCardByString receives string like 'K???' and returns appropriate *Card from Locale
func (l *Locale) FindCardByString(s string) (*Card, error) {
	runes := []rune(s)

	if len(runes) < 2 || len(runes) > 3 {
		return nil, fmt.Errorf("Unable to find card: invalid length of string '%v': %v", s, len(s))
	}

	// FIXME: we should rely on the correct order of cards in a locale
	for _, card := range l.Cards {
		if card.Rank == string(runes[:len(runes)-1]) && card.Suit == string(runes[len(runes)-1:]) {
			return &card, nil
		}
	}

	return nil, fmt.Errorf("Failed to find card `%s` in the locale", s)
}

// Validate Locale
func (l *Locale) Validate() error {
	if errs := validator.Validate(l); errs != nil {
		return errs
	}
	return nil
}

// NewLocale reads file with locale, tries to unmarshall it and on success
// returns *Locale
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

	loc.Core.od = NewOrderedDeck(loc)
	loc.Core.om = NewOriginMatrix(&origin, loc.Core.od)
	loc.Core.hm = NewHumansMatrix(loc.Core.om, loc.Core.od)
	loc.Core.am = NewAngelsMatrix(loc.Core.om, loc.Core.od)
	loc.Core.mm = NewMatrices(loc.Core.om, loc.Core.od)
	loc.Core.cc = NewCyclesTable()

	return loc, nil
}

// MustBuildLocales returns Locales (map[language.Tag]*Locale) or raises panic
// It doesn't make sense to continue without locales
func MustBuildLocales(paths ...string) Locales {
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
