package gosseract

type tesseract0303 struct {
	version string
}

func (t tesseract0303) Version() string {
	return t.version
}
