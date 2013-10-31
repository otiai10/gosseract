package gosseract

import (
)

/**
 * サーバントの言語関係のメソッドを
 * 集約するファイル
 */

func (s *Servant) AvailableLanguages() []string {
  langs := []string{}
  for _, lang := range getAvailableLanguages() {
    langs = append(langs, lang)
  }
  return langs
}
