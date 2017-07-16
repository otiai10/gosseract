package gosseract

import "fmt"
import "os"
import "os/exec"
import "bytes"
import "io/ioutil"

type tesseract0305 struct {
	version        string
	resultFilePath string
	commandPath    string
}

func (t tesseract0305) Version() string {
	return t.version
}

func (t tesseract0305) Execute(params []string) (res string, e error) {
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

	if len(params) > 1 {
		for i := 1; i < len(params); i++ {
			args = append(args, params[i])
		}
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
	var hocr bool
	for _, a := range args {
		if a == "hocr" {
			hocr = true
		}
	}

	if hocr {
		res, e = t.readResult(".hocr")
	} else {
		res, e = t.readResult(outFILEEXTENSION)
	}
	return
}

func (t tesseract0305) readResult(extenstion string) (res string, e error) {
	fpath := t.resultFilePath + extenstion
	file, e := os.OpenFile(fpath, 1, 1)
	if e != nil {
		return
	}
	buffer, _ := ioutil.ReadFile(file.Name())
	res = string(buffer)
	os.Remove(file.Name())
	return
}
