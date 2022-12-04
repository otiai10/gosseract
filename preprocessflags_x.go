package gosseract

// #cgo CXXFLAGS: -std=c++0x
// #cgo LDFLAGS: -llept -ltesseract
// #cgo CPPFLAGS: -Wno-unused-result
//
// #ifdef __APPLE__
//   #cgo LDFLAGS: -L/usr/local/lib -llept -ltesseract
// #endif
import "C"
