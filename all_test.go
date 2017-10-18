package gosseract

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestVersion(t *testing.T) {
	// client, err := NewClient()
	// Expect(t, err).ToBe(nil)
	// Expect(t, client).TypeOf("*gosseract.APIClient")

	version := Version()
	Expect(t, version).ToBe("3.05.00")
}
