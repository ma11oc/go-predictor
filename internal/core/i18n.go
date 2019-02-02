package core

import (
	"fmt"
	"io/ioutil"

	"github.com/go-validator/validator"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	yaml "gopkg.in/yaml.v2"
)

// loadtranslation
//
type Locale struct {
	Lang     string `yaml:"lang" validate:"nonzero"`
	Matrices []struct {
		ID string `yaml:"id"`
	}
	Cards []Card `yaml:"cards" validate:"len=52"`
}

func (l Locale) Validate() error {
	if errs := validator.Validate(l); errs != nil {
		return errs
	}
	return nil
}

func MustLoadLocale(p string) error {
	var content []byte
	var err error

	content, err = ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	locale := &Locale{}

	err = yaml.Unmarshal([]byte(content), &locale)
	if err != nil {
		panic(err)
	}

	if err = locale.Validate(); err != nil {
		panic(fmt.Errorf("Locale %v is invalid: %v\n", p, err))
	}

	lang := language.Make(locale.Lang)

	for _, c := range locale.Cards {
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".title"), c.Title)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.general.keywords"), c.Meanings.General.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.general.description"), c.Meanings.General.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.longterm.keywords"), c.Meanings.Longterm.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.longterm.description"), c.Meanings.Longterm.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.mercury.keywords"), c.Meanings.Mercury.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.mercury.description"), c.Meanings.Mercury.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.venus.keywords"), c.Meanings.Venus.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.venus.description"), c.Meanings.Venus.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.mars.keywords"), c.Meanings.Mars.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.mars.description"), c.Meanings.Mars.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.jupiter.keywords"), c.Meanings.Jupiter.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.jupiter.description"), c.Meanings.Jupiter.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.saturn.keywords"), c.Meanings.Saturn.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.saturn.description"), c.Meanings.Saturn.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.uranus.keywords"), c.Meanings.Uranus.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.uranus.description"), c.Meanings.Uranus.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.neptune.keywords"), c.Meanings.Neptune.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.neptune.description"), c.Meanings.Neptune.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.pluto.keywords"), c.Meanings.Pluto.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.pluto.description"), c.Meanings.Pluto.Description)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.result.keywords"), c.Meanings.Result.Keywords)
		message.SetString(lang, fmt.Sprintf("%v%v", c.ID, ".meanings.result.description"), c.Meanings.Result.Description)
	}
	return nil
}

func LoadLocales(paths ...string) {
	for _, p := range paths {
		MustLoadLocale(p)
	}
}
