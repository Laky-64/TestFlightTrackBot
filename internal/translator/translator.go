package translator

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strings"
)

type Translator struct {
	localizer *i18n.Localizer
}

func New(language string) *Translator {
	return &Translator{
		localizer: i18n.NewLocalizer(bundle, language),
	}
}

func (t *Translator) T(key Key) string {
	return t.TWithData(key, nil)
}

func (t *Translator) TWithData(key Key, data map[string]interface{}) string {
	msg, err := t.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    string(key),
		TemplateData: data,
		DefaultMessage: &i18n.Message{
			ID:    string(key),
			Other: string(key),
		},
	})
	if err != nil {
		return "%" + string(key) + "%"
	}
	return msg
}

func TAll(key Key) []string {
	var translations []string
	for _, lang := range SupportedLanguages() {
		translations = append(translations, New(lang).T(key))
	}
	return translations
}

func TKeys() []Key {
	return availableKeys
}

func SupportedLanguages() []string {
	var langList []string
	for _, tag := range bundle.LanguageTags() {
		langList = append(langList, tag.String())
	}
	return langList
}

func IsSupported(lang string) bool {
	for _, tag := range bundle.LanguageTags() {
		if tag.String() == lang {
			return true
		}
	}
	return false
}

func LangToFlag(lang string) string {
	lang = strings.ReplaceAll(lang, "en", "gb")
	if len(lang) != 2 {
		return ""
	}
	a := []rune(strings.ToUpper(lang))
	const base = 0x1F1E6
	flag := []rune{
		base + (a[0] - 'A'),
		base + (a[1] - 'A'),
	}
	return string(flag)
}
