package gosseract

import "fmt"
import "os/exec"
import "bytes"
import "regexp"
import "io/ioutil"

type tesseractCmd interface {
	Version() string
	Execute(args []string) (string, error)
}

const TESSERACT = "tesseract"
const tmpFILEPREFIX = "gosseract"
const outFILEEXTENSION = ".txt"

func getTesseractCmd() (tess tesseractCmd, e error) {
	commandPath, e := lookPath()
	if e != nil {
		return
	}
	v, e := version()
	if e != nil {
		return
	}
	if regexp.MustCompile("^3.02").Match([]byte(v)) {
		tess = tesseract0302{version: v, commandPath: commandPath}
		return
	}
	if regexp.MustCompile("^3.03").Match([]byte(v)) {
		tess = tesseract0303{version: v, commandPath: commandPath}
		return
	}
	if regexp.MustCompile("^3.04").Match([]byte(v)) {
		tess = tesseract0304{version: v, commandPath: commandPath}
		return
	}
	if regexp.MustCompile("^3.05").Match([]byte(v)) {
		tess = tesseract0305{version: v, commandPath: commandPath}
		return
	}
	e = fmt.Errorf("No tesseract version is found, supporting 3.02~, 3.03~, 3.04~ and 3.05~")
	return
}
func lookPath() (commandPath string, e error) {
	return exec.LookPath(TESSERACT)
}
func version() (v string, e error) {
	v, e = execTesseractCommandWithStderr("--version")
	if e != nil {
		return
	}
	exp := regexp.MustCompile("^tesseract ([0-9\\.]+)")
	matches := exp.FindStringSubmatch(v)
	if len(matches) < 2 {
		e = fmt.Errorf("tesseract version not found: response is `%s`", v)
		return
	}
	v = matches[1]
	return
}
func execTesseractCommandWithStderr(opt string) (res string, e error) {
	cmd := exec.Command(TESSERACT, opt)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if e = cmd.Run(); e != nil {
		return
	}
	res = stdout.String() + stderr.String()
	return
}
func generateTmpFile() (fname string, e error) {
	myTmpDir := "" // TODO: enable to choose optionally
	f, e := ioutil.TempFile(myTmpDir, tmpFILEPREFIX)
	if e != nil {
		return
	}
	fname = f.Name()
	return
}
