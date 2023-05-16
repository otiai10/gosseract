package gosseract

// #cgo CXXFLAGS: -std=c++0x
// #cgo CPPFLAGS: -I/usr/local/include
// #cgo CPPFLAGS: -Wno-unused-result
// #cgo darwin LDFLAGS: -L/usr/local/lib -llept -ltesseract
// #cgo !darwin LDFLAGS: -L/usr/local/lib -lleptonica -ltesseract
import "C"
