package gosseract

import (
  . "os"
  "os/exec"
  "io/ioutil"
)

type AnywayArgs struct {
  SourcePath  string
  Destination string
}
var (
  TMPDIR = "/tmp"
  OUTEXT = ".txt"
  COMMAND = "tesseract"
)

func HelloGosseract() string {
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
  f, _ := OpenFile(fn, 1, 1)
  buf, _ := ioutil.ReadFile(f.Name())
  out = string(buf)

  return out
}
