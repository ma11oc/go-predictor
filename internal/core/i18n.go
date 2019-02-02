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
	validator.SetValidationFunc("langtag", isParsableLanguageTag)
}

// Make sure a value of a field is parsible by language.Parse()
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
type Locale struct {
	Lang     string `yaml:"lang" validate:"nonzero,langtag"`
	Matrices []struct {
		ID string `yaml:"id"`
	}
	Cards []Card `yaml:"cards" validate:"len=52"`
}

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

	locale := &Locale{}

	if err = yaml.Unmarshal([]byte(content), &locale); err != nil {
		return nil, fmt.Errorf("unable to unmarshall locale: %v", err)
	}

	if err = locale.Validate(); err != nil {
		return nil, fmt.Errorf("locale %v is invalid: %v", p, err)
	}

	return locale, nil
}

func MustLoadLocales(paths ...string) map[language.Tag]*Locale {
	var locale *Locale
	var err error

	locales := make(map[language.Tag]*Locale, len(paths))

	for _, p := range paths {
		if locale, err = NewLocale(p); err != nil {
			panic(err)
		}

		// Lang has been validated before, so there is not need to parse it
		lang := language.Make(locale.Lang)

		locales[lang] = locale
	}

	return locales
}
