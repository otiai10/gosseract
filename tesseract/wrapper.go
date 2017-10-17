package tesseract

/*
#if __FreeBSD__ >= 10
#cgo LDFLAGS: -L/usr/local/lib -llept -ltesseract
#else
#cgo LDFLAGS: -llept -ltesseract
#endif

#include "tess.h"
*/
import "C"

// Simple executes tesseract only with source image file path.
func Simple(imgPath string, whitelist string, languages string) string {
	p := C.CString(imgPath)
	w := C.CString(whitelist)
	l := C.CString(languages)

	s := C.simple(p, w, l)
	return C.GoString(s)
}
