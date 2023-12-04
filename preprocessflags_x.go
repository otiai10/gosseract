//go:build !darwin

package gosseract

// #cgo CXXFLAGS: -std=c++0x
// #cgo CPPFLAGS: -I/usr/local/include
// #cgo CPPFLAGS: -Wno-unused-result
// #cgo LDFLAGS: -L/usr/local/lib -lleptonica -ltesseract
import "C"
