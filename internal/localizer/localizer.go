package localizer

import (
	_ "github.com/alextanhongpin/go-text/internal/translations"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Localizer struct {
	ID      string
	printer *message.Printer
}

var locales = []Localizer{
	{
		// Germany
		ID:      "de-de",
		printer: message.NewPrinter(language.MustParse("de-DE")),
	},
	{
		// Switzerland (French speaking)
		ID:      "fr-ch",
		printer: message.NewPrinter(language.MustParse("fr-CH")),
	},
	{
		// United Kingdom
		ID:      "en-gb",
		printer: message.NewPrinter(language.MustParse("en-GB")),
	},
}

func Get(id string) (Localizer, bool) {
	for _, locale := range locales {
		if id == locale.ID {
			return locale, true
		}
	}
	return Localizer{}, false
}

func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
