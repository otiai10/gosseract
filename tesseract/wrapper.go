package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"

// Simple executes tesseract only with source image file path.
func Simple(imgPath string) string {
	p := C.CString(imgPath)
	s := C.simple(p)
	return C.GoString(s)
}
