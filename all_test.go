package gosseract

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestVersion(t *testing.T) {
	version := Version()
	Expect(t, version).ToBe("3.05.00")
}
