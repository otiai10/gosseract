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

// TODO?: Support windows? :(
var (
	TMPDIR  = "/tmp"
	OUTEXT  = ".txt"
	COMMAND = "tesseract"
	VERSION = "0.0.1"
)

func Greeting() string {
	return "Hello,Gosseract!"
}

// `Anyway` provide the way to execute OCR instantly and directly.
func Anyway(args AnywayArgs) string {
	out := ""
	// TODO: DO NOT USE tmp files, using stdin is better
	//	   @see https://code.google.com/p/tesseract-ocr/issues/detail?id=813
	// TODO: Check tesseract-ocr's version?
	if args.Destination == "" {
		args.Destination = genTmpFilePath()
	}
	// Execute the command
	command := exec.Command(COMMAND, args.SourcePath, args.Destination)
	e := command.Run()
	if e != nil {
		panic(e)
	}

	// TODO: DRY
	// Reading output
	// (outputs of `tesseract` automatically be `{second-args}.txt` format)
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
	command.Stderr = &stderr // XXX: Why it's stderr X(
	e := command.Run()
	if e != nil {
		panic(e)
	}
	// ugly
	tesseractInfo := strings.Split(stderr.String(), " ")[1]
	return strings.TrimRight(tesseractInfo, "\n")
}

// Get all available language able to use from `tesseract`
func getAvailableLanguages() []string {
	command := exec.Command(COMMAND, "--list-langs")
	var stderr bytes.Buffer
	command.Stderr = &stderr // XXX: Why it's stderr X(
	e := command.Run()
	if e != nil {
		panic(e)
	}
	langs := strings.Split(stderr.String(), "\n")
	return langs[1 : len(langs)-1]
}

// Capsulize files management.
// Takes path to source file.
// Returns result string.
func execute(source string, args []string) string {
	_args := []string{}
	_args = append(_args, source)

	dest := genTmpFilePath()

	_args = append(_args, dest)
	for _, a := range args {
		_args = append(_args, a)
	}
	_ = _exec(COMMAND, _args)

	// TODO: DRY
	// Reading output
	// (outputs of `tesseract` automatically be `{second-args}.txt` format)
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

// the very general command execution wrapper
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

// Generates tmp filepath
func genTmpFilePath() string {
	id, _ := uuid.NewV4()
	return TMPDIR + "/" + id.String()
}
