package gosseract

// Must execute tesseract-OCR directly by parameter map
func Must(params map[string]string) (out string) {
	client, _ := NewClient()
	out, _ = client.Must(params)
	return
}
