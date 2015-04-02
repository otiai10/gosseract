package gosseract

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func Test_Must(t *testing.T) {
	Expect(t, Must(Params{
		Src: "./.samples/png/sample000.png",
	})).ToBe("01:37:58\n\n")
}

// tesseract ./.samples/png/sample000.png out -l eng ./.samples/option/digest001.txt
func Test_Must_WithDigest(t *testing.T) {
	params := Params{
		Src: "./.samples/png/sample001.png",
	}
	Expect(t, Must(params)).ToBe("03:41:26\n\n")

	// add optional digest
	// params["digest"] = "./.samples/option/digest001.txt"
	params.Whitelist = "IO"
	Expect(t, Must(params)).ToBe("O   I  \n\n")
}

func Test_NewClient(t *testing.T) {
	client, e := NewClient()
	Expect(t, e).ToBe(nil)
	Expect(t, client).TypeOf("*gosseract.Client")
}

func TestClient_Must(t *testing.T) {
	client, _ := NewClient()
	params := map[string]string{}
	_, e := client.Must(params)
	Expect(t, e).Not().ToBe(nil)
}

func TestClient_Src(t *testing.T) {
	client, _ := NewClient()
	out, e := client.Src("./.samples/png/sample000.png").Out()
	Expect(t, e).ToBe(nil)
	Expect(t, out).ToBe("01:37:58\n\n")
}

func TestClient_Image(t *testing.T) {
	client, _ := NewClient()
	img := fixImage("./.samples/png/sample000.png")
	out, e := client.Image(img).Out()
	Expect(t, e).ToBe(nil)
	Expect(t, out).ToBe("01:37:58\n\n")
}

func TestClient_Digest(t *testing.T) {
	client, _ := NewClient()
	img := fixImage("./.samples/png/sample001.png")
	out, e := client.Image(img).Out()
	Expect(t, e).ToBe(nil)
	Expect(t, out).ToBe("03:41:26\n\n")

	out, e = client.Digest("./.samples/option/digest001.txt").Image(img).Out()
	Expect(t, e).ToBe(nil)
	Expect(t, out).ToBe("O   I  \n\n")
}

func fixImage(fpath string) image.Image {
	f, _ := os.Open(fpath)
	buf, _ := ioutil.ReadFile(f.Name())
	img, _ := png.Decode(bytes.NewReader(buf))
	return img
}

func TestClient_Out(t *testing.T) {
	client, _ := NewClient()
	_, e := client.Out()
	Expect(t, e.Error()).ToBe("Source is not set")
}
