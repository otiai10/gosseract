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
  Source  Source
  Lang    Lang
  Options Options
}
type Source struct {
  FilePath string
  // 画像形式とかくるんだろうな今後
}
type Lang struct {
  Value      string
  Availables []string
}
type Options struct {
  UseFile   bool
  FilePath  string
  WhiteList string
}
type VersionInfo struct {
  TesseractVersion string
  GosseractVersion string
}

func SummonServant() Servant {
  lang := Lang{}
  lang.init()
  opts := Options{}
  opts.init()
  return Servant{
    Lang:    lang,
    Options: opts,
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

func (s *Servant) Eat(filepath string) *Servant {
  // TODO: ファイル存在チェック
  s.Source.FilePath = filepath
  return s
}

func (s *Servant) Out() (string, /* TODO#1: Error */bool) {
  result := execute(s.Source.FilePath, s.buildArguments())
  return result, false
}

func (s *Servant) buildArguments() []string {
  var args []string
  args = append(args, "-l", s.Lang.Value)
  if s.Options.UseFile {
    args = append(args, s.Options.FilePath)
  }
  return args
}
