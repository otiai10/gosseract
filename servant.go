package gosseract

import (
)

/**
 * Gosseract Servant は
 * tessearctのバージョン取得とか
 * 利用可能言語取得、設定とか
 * ヒントとかを設定できるのﾃﾞｪｽ!!
 */
type Servant struct {
  Lang Lang
}
type Lang struct {
  Value string
}
type VersionInfo struct {
  TesseractVersion string
  GosseractVersion string
}

func SummonServant() Servant {
  lang := Lang{
    Value: "eng",// "eng" in default
  }
  return Servant{
    Lang: lang,
  }
}

func (s *Servant) Greeting() string {
  return "Hi, I'm gosseract-ocr servant!"
}

func (s *Servant) Info() VersionInfo {
  tessVersion := getTesseractVersion()
  info := VersionInfo{
    TesseractVersion: tessVersion,
    GosseractVersion: VERSION,
  }
  return info
}
