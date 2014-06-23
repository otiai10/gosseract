package gosseract

import "fmt"
import "os/exec"
import "bytes"
import "regexp"

type tesseractCmd interface {
	Version() string
}

func GetTesseractCmd() (tess tesseractCmd, e error) {
	v, e := version()
	if e != nil {
		return
	}
	if regexp.MustCompile("3.02").Match([]byte(v)) {
		tess = tesseract0302{v}
		return
	}
	if regexp.MustCompile("3.03").Match([]byte(v)) {
		tess = tesseract0303{v}
		return
	}
	e = fmt.Errorf("No tesseract version is found, supporting 3.02~ and 3.03~.")
	return
}
func version() (v string, e error) {
	v, e = execTesseractCommandWithStderr("--version")
	if e != nil {
		return
	}
	exp := regexp.MustCompile("^tesseract ([0-9\\.]+)")
	matches := exp.FindStringSubmatch(v)
	if len(matches) < 2 {
		e = fmt.Errorf("tesseract version not found: response is `%s`.", v)
	}
	v = matches[1]
	return
}
func execTesseractCommandWithStderr(opt string) (res string, e error) {
	cmd := exec.Command("tesseract", opt)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if e = cmd.Run(); e != nil {
		return
	}
	res = stderr.String()
	return
}
