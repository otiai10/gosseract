//go:build windows

package gosseract

// #cgo CFLAGS: -Wno-unused-result
// #cgo LDFLAGS: -ltesseract -lleptonica
import "C"
