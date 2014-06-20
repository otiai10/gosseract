package gosseract_test

import "github.com/otiai10/gosseract"
import . "github.com/otiai10/mint"
import "testing"

func TestNewClient(t *testing.T) {
	Expect(t, gosseract.Greet()).ToBe("Hello,Gosseract.")
}
