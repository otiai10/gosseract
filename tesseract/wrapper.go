package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"

// Simple executes tesseract only with source image file path.
func Simple(imgPath string, whitelist string,languages string) string {
	p := C.CString(imgPath)
	w := C.CString(whitelist)
	l := C.CString(languages)

	s := C.simple(p, w,l)
	return C.GoString(s)
}
