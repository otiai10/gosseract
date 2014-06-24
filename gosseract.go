package gosseract

// Must execute tesseract-OCR directly by parameter map
func Must(params map[string]string) (out string) {
	client, e := NewClient()
	if e != nil {
		return
	}
	out, _ = client.Must(params)
	return
}
