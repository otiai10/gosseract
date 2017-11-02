package gosseract

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestVersion(t *testing.T) {
	version := Version()
	Expect(t, version).Match("[0-9]{1}.[0-9]{2}(.[0-9a-z]*)")
}

func TestNewClient(t *testing.T) {
	client := NewClient()
	defer client.Close()

	Expect(t, client).TypeOf("*gosseract.Client")
}

func TestClient_SetImage(t *testing.T) {
	client := NewClient()
	defer client.Close()

	client.Trim = true
	client.SetImage("./test/data/001-gosseract.png")

	client.SetPageSegMode(PSM_SINGLE_BLOCK)

	text, err := client.Text()
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("otiai10 / gosseract")

	client.Languages = []string{"eng"}
	client.SetWhitelist("aegitosr10/")
	text, err = client.Text()
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("otiai10 / gosseraet")
}
