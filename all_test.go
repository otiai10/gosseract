package gosseract_test

import "github.com/otiai10/gosseract"
import . "github.com/otiai10/mint"
import "testing"
import "os"
import "bytes"
import "image"
import "image/png"
import "io/ioutil"

func Test_Must(t *testing.T) {
	params := map[string]string{
		"src": "./.samples/png/sample000.png",
	}
	Expect(t, gosseract.Must(params)).ToBe("01:37:58\n\n")
}

// tesseract ./.samples/png/sample000.png out -l eng ./.samples/option/digest001.txt
func Test_Must_WithDigest(t *testing.T) {
	params := map[string]string{
		"src": "./.samples/png/sample001.png",
	}
	Expect(t, gosseract.Must(params)).ToBe("03:41:26\n\n")

	// add optional digest
	params["digest"] = "./.samples/option/digest001.txt"
	Expect(t, gosseract.Must(params)).ToBe("O   I  \n\n")
}

func Test_NewClient(t *testing.T) {
	client, e := gosseract.NewClient()
	Expect(t, e).ToBe(nil)
	Expect(t, client).TypeOf("*gosseract.Client")
}

func TestClient_Must(t *testing.T) {
	client, _ := gosseract.NewClient()
	params := map[string]string{}
	_, e := client.Must(params)
	Expect(t, e).Not().ToBe(nil)
}

func TestClient_Src(t *testing.T) {
	client, _ := gosseract.NewClient()
	out, e := client.Src("./.samples/png/sample000.png").Out()
	Expect(t, e).ToBe(nil)
	Expect(t, out).ToBe("01:37:58\n\n")
}

func TestClient_Image(t *testing.T) {
	client, _ := gosseract.NewClient()
	img := fixImage("./.samples/png/sample000.png")
	out, e := client.Image(img).Out()
	Expect(t, e).ToBe(nil)
	Expect(t, out).ToBe("01:37:58\n\n")
}
func fixImage(fpath string) image.Image {
	f, _ := os.Open(fpath)
	buf, _ := ioutil.ReadFile(f.Name())
	img, _ := png.Decode(bytes.NewReader(buf))
	return img
}

func TestClient_Out(t *testing.T) {
	client, _ := gosseract.NewClient()
	_, e := client.Out()
	Expect(t, e.Error()).ToBe("Source is not set")
}
