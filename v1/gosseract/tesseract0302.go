package gosseract

import "fmt"
import "os"
import "os/exec"
import "bytes"
import "io/ioutil"

type tesseract0302 struct {
	version        string
	resultFilePath string
	commandPath    string
}

func (t tesseract0302) Version() string {
	return t.version
}
func (t tesseract0302) Execute(params []string) (res string, e error) {

	// command args
	var args []string
	// Register source file
	args = append(args, params[0])
	// generate result file path
	t.resultFilePath, e = generateTmpFile()
	if e != nil {
		return
	}
	// Register result file
	args = append(args, t.resultFilePath)
	// Register digest file
	if len(params) > 1 {
		args = append(args, params[1])
	}

	// prepare command
	cmd := exec.Command(TESSERACT, args...)
	// execute
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if e = cmd.Run(); e != nil {
		e = fmt.Errorf(stderr.String())
		return
	}
	// read result
	res, e = t.readResult()
	os.Remove(t.resultFilePath)
	return
}
func (t tesseract0302) readResult() (res string, e error) {
	fpath := t.resultFilePath + outFILEEXTENSION
	file, e := os.OpenFile(fpath, 1, 1)
	if e != nil {
		return
	}
	buffer, _ := ioutil.ReadFile(file.Name())
	res = string(buffer)
	os.Remove(file.Name())
	return
}
