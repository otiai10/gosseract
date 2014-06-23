package gosseract

type tesseract0303 struct {
	version string
}

func (t tesseract0303) Version() string {
	return t.version
}
func (t tesseract0303) Execute(args []string) (res string, e error) {
	res = "tesseract0303"
	return
}
