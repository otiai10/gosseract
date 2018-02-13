package gosseract

import (
	"strings"
	"testing"

	"golang.org/x/net/html"

	. "github.com/otiai10/mint"
)

func TestVersion(t *testing.T) {
	version := Version()
	Expect(t, version).Match("[0-9]{1}.[0-9]{2}(.[0-9a-z]*)?")
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

func TestClient_SetLanguage(t *testing.T) {
	client := NewClient()
	defer client.Close()
	client.SetLanguage("deu")
	client.SetImage("./test/data/001-gosseract.png")
	_, err := client.Text()
	Expect(t, err).Not().ToBe(nil)
}

func TestClient_ConfigFilePath(t *testing.T) {

	client := NewClient()
	defer client.Close()

	err := client.SetConfigFile("./test/config/01.config")
	Expect(t, err).ToBe(nil)
	client.SetImage("./test/data/001-gosseract.png")
	text, err := client.Text()
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("otiai10 l gosseract")

	When(t, "the config file is not found", func(t *testing.T) {
		err := client.SetConfigFile("./test/config/not-existing")
		Expect(t, err).Not().ToBe(nil)
	})

	When(t, "the config file path is a directory", func(t *testing.T) {
		err := client.SetConfigFile("./test/config/02.config")
		Expect(t, err).Not().ToBe(nil)
	})

}

func TestClient_HTML(t *testing.T) {
	client := NewClient()
	defer client.Close()
	client.SetImage("./test/data/001-gosseract.png")
	client.SetWhitelist("otiai10/gosseract")
	out, err := client.HOCRText()
	Expect(t, err).ToBe(nil)

	tokenizer := html.NewTokenizer(strings.NewReader(out))

	texts := []string{}
	for ttype := tokenizer.Next(); ttype != html.ErrorToken; ttype = tokenizer.Next() {
		token := tokenizer.Token()
		if token.Type == html.TextToken && strings.TrimSpace(token.Data) != "" {
			texts = append(texts, strings.Trim(token.Data, "\n"))
		}
	}
	Expect(t, texts).ToBe([]string{"otiai10", "/", "gosseract"})

	When(t, "only invalid languages are given", func(t *testing.T) {
		client := NewClient()
		defer client.Close()
		client.SetLanguage("foo")
		client.SetImage("./test/data/001-gosseract.png")
		_, err := client.HOCRText()
		Expect(t, err).Not().ToBe(nil)
	})
	When(t, "undefined key-value is tried to be set", func(t *testing.T) {
		client := NewClient()
		defer client.Close()
		client.SetVariable("foobar", "hoge")
		client.SetImage("./test/data/001-gosseract.png")
		_, err := client.HOCRText()
		Expect(t, err).Not().ToBe(nil)
	})
}
