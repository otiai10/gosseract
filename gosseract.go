package gosseract

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/nu7hatch/gouuid"
)

type AnywayArgs struct {
	SourcePath  string
	Destination string
}

var (
	TMPDIR  = "/tmp"
	OUTEXT  = ".txt"
	COMMAND = "tesseract"
	VERSION = "0.0.1"
)

func Greeting() string {
	return "Hello,Gosseract!"
}

// func `gosseract.Anyway` can OCR from multi args
func Anyway(args AnywayArgs) string {
	// 最終的な返り値
	out := ""
	// tesseractが標準出力に対応してるハズ
	// tesseractのバージョンを見るようなメソッドを用意しないとアカンなこれ
	if args.Destination == "" {
		args.Destination = genTmpFilePath()
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

	_ = os.Remove(f.Name())

	return out
}

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
	return langs[1 : len(langs)-1]
}

/**
 * ソースファイルパスと
 * オプションアーギュメントのスライスを受け取り
 * OCRしたものを返す
 * ファイル操作などを隠蔽する
 */
func execute(source string, args []string) string {
	_args := []string{}
	_args = append(_args, source)

	dest := genTmpFilePath()

	_args = append(_args, dest)
	for _, a := range args {
		_args = append(_args, a)
	}
	_ = _exec(COMMAND, _args)

	// 出力を読む
	// tesseractの出力はコマンドラインの第二引数に.txtを付けたものに置かれる
	// TODO4: DRY
	fn := dest + OUTEXT

	f, _ := os.OpenFile(fn, 1, 1)
	buf, _ := ioutil.ReadFile(f.Name())
	out := string(buf)

	_ = os.Remove(fn)

	return out
}

func tesseractInstalled() bool {
	found := _exec("which", []string{"tesseract"})

	return found != ""
}

/**
 * 汎用: コマンドを実行する
 */
func _exec(command string, args []string) string {
	cmd := exec.Command(command, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	_ = cmd.Run()
	if stdout.String() != "" {
		return stdout.String()
	}
	return stderr.String()
}

func genTmpFilePath() string {
	id, _ := uuid.NewV4()
	return TMPDIR + "/" + id.String()
}
