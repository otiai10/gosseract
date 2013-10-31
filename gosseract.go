package gosseract

import (
  "os"
  "os/exec"
  "io/ioutil"
  "bytes"
  "strings"
)
/* TODO#1: Error or nil っていう返し方無いか
type Error struct {
  Message string
}
TODO#1が解決されるまでコメントアウト */

type AnywayArgs struct {
  SourcePath  string
  Destination string
}
var (
  TMPDIR = "/tmp"
  OUTEXT = ".txt"
  COMMAND = "tesseract"
  VERSION = "0.0.1"
)

func Greeting() string {
  return "Hello,Gosseract!"
}

/**
 * とにかくパラメータ喰わせて一発でOCRしたい場合の
 * コマンドラッパー
 */
func Anyway(args AnywayArgs) string {
  // 最終的な返り値
  out := ""

  // 引数で行き先を指定されない場合は
  // (とりあえず) `/tmp/anyway,txt`に置く
  // tesseractが標準出力に対応してるハズ
  // tesseractのバージョンを見るようなメソッドを用意しないとアカンなこれ
  if args.Destination == "" {
    args.Destination = TMPDIR + "/anyway"
  }

  // tesseractコマンドを実行
  command := exec.Command(COMMAND, args.SourcePath, args.Destination)
  e := command.Run()
  if e != nil {
    panic(e)
  }

  // 出力を読む
  // tesseractの出力はコマンドラインの第二引数に.txtを付けたものに置かれる
  fn := args.Destination + OUTEXT
  f, _ := os.OpenFile(fn, 1, 1)
  buf, _ := ioutil.ReadFile(f.Name())
  out = string(buf)

  return out
}

/**
 * privateなメソッドってどうやって定義するんだ？
 */
func getTesseractVersion() string {
  command := exec.Command(COMMAND, "--version")
  var stderr bytes.Buffer
  command.Stderr = &stderr //謎に標準エラーで来るw
  e := command.Run()
  if e != nil {
    panic(e)
  }
  // なんかクズい
  tesseractInfo := strings.Split(stderr.String(), " ")[1]
  return strings.TrimRight(tesseractInfo, "\n")
}
/**
 * 利用可能な言語の一覧を取得する
 */
func getAvailableLanguages() []string {
  command := exec.Command(COMMAND, "--list-langs")
  var stderr bytes.Buffer
  command.Stderr = &stderr //謎に標準エラーで来るw
  e := command.Run()
  if e != nil {
    panic(e)
  }
  langs := strings.Split(stderr.String(), "\n")
  langs = langs[1:len(langs) - 1]
  return langs
}
