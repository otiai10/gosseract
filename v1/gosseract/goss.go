package gosseract

import "github.com/otiai10/gosseract/v1/gosseract/tesseract"

// Params is parameters for gosseract.Must.
type Params struct {
	Src       string // source image file path
	Whitelist string // tessedit_char_whitelist
	Languages string
}

// Must execute tesseract-OCR directly by parameter map
func Must(params Params) (out string) {
	return tesseract.Simple(params.Src, params.Whitelist,params.Languages)
}
