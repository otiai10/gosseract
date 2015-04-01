package gosseract

import "github.com/otiai10/gosseract/tesseract"

// Params is parameters for gosseract.Must.
type Params struct {
	Src string
}

// Must execute tesseract-OCR directly by parameter map
func Must(params Params) (out string) {
	return tesseract.Simple(params.Src)
}
