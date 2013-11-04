/**
 * サーバントの言語関係のメソッドを
 * 集約するファイル
 */
package gosseract

import (
  "errors"
)

func (l *Lang) Available() []string {
  return l.Availables
}
func (l *Lang) Have(key string) bool {
  for _, language := range l.Availables {
    if language == key {
      return true
    }
  }
  return false
}
func (l *Lang) Is() string {
  return l.Value
}
func (l *Lang) Use(key string) error {
  if l.Have(key) {
    l.Value = key
    return nil
  }
  return errors.New("とりあえず生のエラー返します")
}

func (l *Lang) init() *Lang {
  l.Value = "eng";// "eng" in default
  l.Availables = getAvailables();
  return l
}

func getAvailables() []string {
  langs := []string{}
  for _, lang := range getAvailableLanguages() {
    langs = append(langs, lang)
  }
  return langs
}
