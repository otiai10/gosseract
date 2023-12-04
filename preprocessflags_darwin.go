package gosseract

// #cgo CXXFLAGS: -std=c++0x
// #cgo CPPFLAGS: -I/opt/homebrew/include -I/usr/local/include
// #cgo CPPFLAGS: -Wno-unused-result
// #cgo LDFLAGS: -L/opt/homebrew/lib -L/usr/local/lib -lleptonica -ltesseract
import "C"
