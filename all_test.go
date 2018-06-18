package gosseract

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"

	. "github.com/otiai10/mint"
)

func TestVersion(t *testing.T) {
	version := Version()
	Expect(t, version).Match("[0-9]{1}.[0-9]{2}(.[0-9a-z]*)?")
}

func TestClearPersistentCache(t *testing.T) {
	client := NewClient()
	defer client.Close()
	client.init()
	ClearPersistentCache()
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
	client.SetImage("./test/data/001-helloworld.png")

	client.SetPageSegMode(PSM_SINGLE_BLOCK)

	text, err := client.Text()
	if client.pixImage == nil {
		t.Errorf("could not set image")
	}
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("Hello, World!")

	err = client.SetImage("./test/data/001-helloworld.png")
	Expect(t, err).ToBe(nil)

	err = client.SetImage("")
	Expect(t, err).Not().ToBe(nil)

	err = client.SetImage("somewhere/fake/fakeimage.png")
	Expect(t, err).Not().ToBe(nil)

	_, err = client.Text()
	Expect(t, err).ToBe(nil)

	Because(t, "api must be initialized beforehand", func(t *testing.T) {
		client := &Client{}
		err := client.SetImage("./test/data/001-helloworld.png")
		Expect(t, err).Not().ToBe(nil)
	})
}

func TestClient_SetImageFromBytes(t *testing.T) {
	client := NewClient()
	defer client.Close()

	content, err := ioutil.ReadFile("./test/data/001-helloworld.png")
	if err != nil {
		t.Fatalf("could not read test file")
	}

	client.Trim = true
	client.SetImageFromBytes(content)

	client.SetPageSegMode(PSM_SINGLE_BLOCK)

	text, err := client.Text()
	if client.pixImage == nil {
		t.Errorf("could not set image")
	}
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("Hello, World!")
	err = client.SetImageFromBytes(content)
	Expect(t, err).ToBe(nil)

	err = client.SetImageFromBytes(nil)
	Expect(t, err).Not().ToBe(nil)

	Because(t, "api must be initialized beforehand", func(t *testing.T) {
		client := &Client{}
		err := client.SetImageFromBytes(content)
		Expect(t, err).Not().ToBe(nil)
	})
}

func TestClient_SetWhitelist(t *testing.T) {

	if os.Getenv("TESS_LSTM") == "1" {
		t.Skip("Whitelist with LSTM is not working for now. Please check https://github.com/tesseract-ocr/tesseract/issues/751")
	}

	client := NewClient()
	defer client.Close()

	client.Trim = true
	client.SetImage("./test/data/001-helloworld.png")
	client.Languages = []string{"eng"}
	client.SetWhitelist("HeloWrd,")
	text, err := client.Text()
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("Hello, Worldl")
}

func TestClient_SetBlacklist(t *testing.T) {

	if os.Getenv("TESS_LSTM") == "1" {
		t.Skip("Blacklist with LSTM is not working for now. Please check https://github.com/tesseract-ocr/tesseract/issues/751")
	}

	client := NewClient()
	defer client.Close()

	client.Trim = true
	err := client.SetImage("./test/data/001-helloworld.png")
	Expect(t, err).ToBe(nil)
	client.Languages = []string{"eng"}
	err = client.SetBlacklist("l")
	Expect(t, err).ToBe(nil)
	text, err := client.Text()
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("He110, WorId!")
}

func TestClient_SetLanguage(t *testing.T) {
	client := NewClient()
	defer client.Close()
	err := client.SetLanguage("deu")
	Expect(t, err).ToBe(nil)
	err = client.SetLanguage()
	Expect(t, err).Not().ToBe(nil)
	client.SetImage("./test/data/001-helloworld.png")
	_, err = client.Text()
	Expect(t, err).Not().ToBe(nil)
}

func TestClient_ConfigFilePath(t *testing.T) {

	if os.Getenv("TESS_LSTM") == "1" {
		t.Skip("Whitelist with LSTM is not working for now. Please check https://github.com/tesseract-ocr/tesseract/issues/751")
	}

	client := NewClient()
	defer client.Close()

	err := client.SetConfigFile("./test/config/01.config")
	Expect(t, err).ToBe(nil)
	client.SetImage("./test/data/001-helloworld.png")
	text, err := client.Text()
	Expect(t, err).ToBe(nil)

	Expect(t, text).ToBe("H    W   ")

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
	client.SetImage("./test/data/001-helloworld.png")
	client.SetWhitelist("Hello,World!")
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
	Expect(t, texts).ToBe([]string{"Hello,", "World!"})

	When(t, "only invalid languages are given", func(t *testing.T) {
		client := NewClient()
		defer client.Close()
		client.SetLanguage("foo")
		client.SetImage("./test/data/001-helloworld.png")
		_, err := client.HOCRText()
		Expect(t, err).Not().ToBe(nil)
	})
	Because(t, "unknown key is validated when `init` is called", func(t *testing.T) {
		client := NewClient()
		defer client.Close()
		err := client.SetVariable("foobar", "hoge")
		Expect(t, err).ToBe(nil)
		client.SetImage("./test/data/001-helloworld.png")
		_, err = client.Text()
		Expect(t, err).Not().ToBe(nil)
	})
}
