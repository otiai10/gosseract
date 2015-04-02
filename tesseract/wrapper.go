package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"

// Simple executes tesseract only with source image file path.
func Simple(imgPath string, whitelist string) string {
	p := C.CString(imgPath)
	w := C.CString(whitelist)
	s := C.simple(p, w)
	return C.GoString(s)
}
