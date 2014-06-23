package gosseract

import "fmt"
import "os"
import "os/exec"
import "bytes"
import "io/ioutil"

type tesseract0302 struct {
	version        string
	resultFilePath string
}

func (t tesseract0302) Version() string {
	return t.version
}
func (t tesseract0302) Execute(args []string) (res string, e error) {
	// generate result file path
	t.resultFilePath, e = generateTmpFile()
	if e != nil {
		return
	}
	// bind it to args
	args = append(args, t.resultFilePath)
	// prepare command
	cmd := exec.Command("tesseract", args...)
	// execute
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if e = cmd.Run(); e != nil {
		e = fmt.Errorf(stderr.String())
		return
	}
	// read result
	res, e = t.readResult()
	return
}
func (t tesseract0302) readResult() (res string, e error) {
	fpath := t.resultFilePath + OUT_FILE_EXT
	file, e := os.OpenFile(fpath, 1, 1)
	if e != nil {
		return
	}
	buffer, _ := ioutil.ReadFile(file.Name())
	res = string(buffer)
	os.Remove(file.Name())
	return
}
