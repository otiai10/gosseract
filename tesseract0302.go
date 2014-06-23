package gosseract

type tesseract0302 struct {
	version string
}

func (t tesseract0302) Version() string {
	return t.version
}
