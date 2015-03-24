package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"

func Hoge() {
	C.hoge()
}

func Do(imgPath string) string {
	p := C.CString(imgPath)
	s := C.fuga(p)
	return C.GoString(s)
}
