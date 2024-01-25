package gosseract

import (
	"encoding/xml"
	"image"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	beforeTest()
	code := m.Run()
	os.Exit(code)
}

func beforeTest() {
	if strings.HasPrefix(Version(), "4.0") {
		os.Setenv("TESS_LSTM_DISABLED", "1")
	}
	switch os.Getenv("TESTCASE") {
	case "archlinux", "centos", "debian", "fedora", "mingw":
		os.Setenv("TESS_BOX_DISABLED", "1")
	}
}

func TestVersion(t *testing.T) {
	version := Version()
	Expect(t, version).Match("[0-9]{1}.[0-9]{1,2}(.[0-9a-z_-]*)?")
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

func TestDoubleClose(t *testing.T) {
	client := NewClient()
	client.Close()
	client.Close()
}

func TestClient_SetTessdataPrefix(t *testing.T) {
	client := NewClient()
	defer client.Close()

	cwd, err := os.Getwd()
	Expect(t, err).ToBe(nil)
	testModelDir := filepath.Join(cwd, "test-model", "tessdata")

	err = os.MkdirAll(testModelDir, 0770)
	Expect(t, err).ToBe(nil)
	defer os.RemoveAll(filepath.Dir(testModelDir))

	src, err := os.Open(filepath.Join(getDataPath(), "eng.traineddata"))
	Expect(t, err).ToBe(nil)
	defer src.Close()

	dst, err := os.Create(filepath.Join(testModelDir, "eng.traineddata"))
	Expect(t, err).ToBe(nil)

	_, err = io.Copy(dst, src)
	Expect(t, err).ToBe(nil)
	dst.Close()

	client.Trim = true
	client.SetImage("./test/data/001-helloworld.png")
	client.SetLanguage("eng")
	client.SetTessdataPrefix(testModelDir)

	text, err := client.Text()
	Expect(t, err).ToBe(nil)
	Expect(t, text).ToBe("Hello, World!")
}

func TestClient_Version(t *testing.T) {
	client := NewClient()
	defer client.Close()
	version := client.Version()
	Expect(t, version).Match("[0-9]{1}.[0-9]{1,2}(.[0-9a-z_-]*)?")
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

	if os.Getenv("TESS_LSTM_DISABLED") == "1" {
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

	// Expect(t, text).ToBe("Hello, Worldl")
	Expect(t, text).Match("Hello, ?Worldl?")
}

func TestClient_SetBlacklist(t *testing.T) {

	if os.Getenv("TESS_LSTM_DISABLED") == "1" {
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
	Expect(t, text).Match("He(110|tto|o), Wor(I|t)?d!")
}

func TestClient_SetLanguage(t *testing.T) {
	client := NewClient()
	defer client.Close()
	err := client.SetLanguage("undefined-language")
	Expect(t, err).ToBe(nil)
	err = client.SetLanguage()
	Expect(t, err).Not().ToBe(nil)
	client.SetImage("./test/data/001-helloworld.png")
	_, err = client.Text()
	Expect(t, err).Not().ToBe(nil)
	if os.Getenv("GOSSERACT_CPPSTDERR_NOT_CAPTURED") != "1" {
		Expect(t, err).Match("Failed loading language 'undefined-language'")
	}
}

func TestClient_ConfigFilePath(t *testing.T) {

	if os.Getenv("TESS_LSTM_DISABLED") == "1" {
		t.Skip("Whitelist with LSTM is not working for now. Please check https://github.com/tesseract-ocr/tesseract/issues/751")
	}

	client := NewClient()
	defer client.Close()

	err := client.SetConfigFile("./test/config/01.config")
	Expect(t, err).ToBe(nil)
	client.SetImage("./test/data/001-helloworld.png")
	text, err := client.Text()
	Expect(t, err).ToBe(nil)

	Expect(t, text).Match("H *W *")

	When(t, "the config file is not found", func(t *testing.T) {
		err := client.SetConfigFile("./test/config/not-existing")
		Expect(t, err).Not().ToBe(nil)
	})

	When(t, "the config file path is a directory", func(t *testing.T) {
		err := client.SetConfigFile("./test/config/02.config")
		Expect(t, err).Not().ToBe(nil)
	})

}

func TestClientBoundingBox(t *testing.T) {

	if os.Getenv("TESS_BOX_DISABLED") == "1" {
		t.Skip()
	}

	client := NewClient()
	defer client.Close()
	client.SetImage("./test/data/001-helloworld.png")
	client.SetWhitelist("Hello,World!")
	boxes, err := client.GetBoundingBoxes(RIL_WORD)
	Expect(t, err).ToBe(nil)

	Because(t, "api must be initialized beforehand", func(t *testing.T) {
		client := &Client{}
		_, err := client.GetBoundingBoxes(RIL_WORD)
		Expect(t, err).Not().ToBe(nil)
	})

	words := []string{"Hello,World!"}
	coords := []image.Rectangle{
		image.Rect(74, 64, 1099, 190),
	}

	if os.Getenv("TESS_LSTM_DISABLED") == "1" {
		t.Skip()
	}
	for i, box := range boxes {
		Expect(t, box.Word).ToBe(words[i])
		Expect(t, box.Box).ToBe(coords[i])
	}
}

func TestClient_HTML(t *testing.T) {

	if os.Getenv("TESS_BOX_DISABLED") == "1" {
		t.Skip()
	}

	client := NewClient()
	defer client.Close()
	client.SetImage("./test/data/001-helloworld.png")
	client.SetWhitelist("Hello,World!")
	out, err := client.HOCRText()
	Expect(t, err).ToBe(nil)

	page := new(Page)
	err = xml.Unmarshal([]byte(out), page)
	Expect(t, err).ToBe(nil)
	Expect(t, len(page.Content.Par.Lines)).ToBe(1)

	if os.Getenv("TESS_LSTM_DISABLED") != "1" {
		Expect(t, len(page.Content.Par.Lines[0].Words)).ToBe(1)
		Expect(t, page.Content.Par.Lines[0].Words[0].Characters).ToBe("Hello,World!")
	}

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

func TestGetAvailableLangs(t *testing.T) {
	t.Skip("TODO")
	// langs, err := GetAvailableLanguages()
	// Expect(t, err).ToBe(nil)
	// Expect(t, len(langs)).ToBe(1) // eng only
}
