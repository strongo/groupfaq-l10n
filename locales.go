package trans

import (
	"fmt"
	"github.com/strongo/app"
	"strings"
)

// GroupFaqLocales - defines locales specific to GroupFAQ app
type GroupFaqLocales struct {
}

// GetLocaleByCode5 - get locale by code
func (GroupFaqLocales) GetLocaleByCode5(code5 string) (locale strongo.Locale, err error) {
	var ok bool
	if locale, ok = SupportedLocalesByCode5[code5]; !ok {
		err = fmt.Errorf("Locale not found by code5: %v", code5)
	}
	return locale, err
}

// SupportedLocalesByCode5 - supported locales by code 5
var SupportedLocales []strongo.Locale = []strongo.Locale{
	strongo.LocaleEnUS,
	strongo.LocaleRuRu,
}

// SupportedLocales  - supported locales
var SupportedLocalesByCode5 = make(map[string]strongo.Locale, len(SupportedLocales))

func init() {
	for _, locale := range SupportedLocales {
		SupportedLocalesByCode5[locale.Code5] = locale
	}
}

// ChooseLocaleIcon - locale icons
var ChooseLocaleIcon = strings.Join([]string{
	strongo.LocaleEnUS.FlagIcon,
	strongo.LocaleRuRu.FlagIcon,
}, " ")
