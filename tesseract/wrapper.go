package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"
import "unsafe"

func Hoge() {
	C.hoge()
}

func Do(imgPath string) string {
	p := C.CString(imgPath)
	s := C.fuga(p)
	defer C.free(unsafe.Pointer(p))
	return C.GoString(s)
}
