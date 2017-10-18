package gosseract

// #cgo LDFLAGS: -llept -ltesseract
// #include <stdlib.h>
// #include "tessbridge.h"
import "C"

// Version returns the version of Tesseract-OCR
func Version() string {
	api := C.Init()
	defer C.Free(api)
	version := C.Version(api)
	return C.GoString(version)
}
