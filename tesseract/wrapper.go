package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"

func Hoge() {
    C.hoge()
}
