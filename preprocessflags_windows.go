//go:build windows

package gosseract

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: tesseract lept
import "C"
