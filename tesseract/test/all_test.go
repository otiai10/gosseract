package tesseract_test

import (
	"testing"

	"github.com/otiai10/gosseract/tesseract"
	. "github.com/otiai10/mint"
)

func TestDo(t *testing.T) {
	Expect(t, tesseract.Simple("hoge.png")).ToBe("otiai10 / gosseract\n\n")
	Expect(t, tesseract.Simple("sample.png")).ToBe("2,464 total\n\n")
}
