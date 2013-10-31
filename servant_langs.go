/**
 * サーバントの言語関係のメソッドを
 * 集約するファイル
 */
package gosseract

import (
)

func (l *Lang) Available() []string {
  return l.Availables
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
