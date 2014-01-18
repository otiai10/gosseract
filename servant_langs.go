package gosseract

import (
	"errors"
)

func (s *Servant) LangAvailable() []string {
	return s.lang.Availables
}
func (s *Servant) LangHave(key string) bool {
	for _, language := range s.lang.Availables {
		if language == key {
			return true
		}
	}
	return false
}
func (s *Servant) LangIs() string {
	return s.lang.Value
}
func (s *Servant) LangUse(key string) error {
	if s.LangHave(key) {
		s.lang.Value = key
		return nil
	}
	return errors.New("Language `" + key + "` is not available.")
}

func (l *lang) init() *lang {
	l.Value = "eng" // "eng" in default
	l.Availables = getAvailables()
	return l
}

func getAvailables() []string {
	langs := []string{}
	for _, lang := range getAvailableLanguages() {
		langs = append(langs, lang)
	}
	return langs
}
