package gosseract_test

import "github.com/otiai10/gosseract"
import . "github.com/otiai10/mint"
import "testing"
import "os"
import "bytes"
import "image"
import "image/png"
import "image/jpeg"
import "io/ioutil"
import "net/http"
import "strings"
import "fmt"

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
func TestClient_Image_by_jpg(t *testing.T) {
	client, _ := gosseract.NewClient()
	img := fixImage("./.samples/jpg/sample001.jpeg", "jpeg")
	out, e := client.Image(img).Out()
	Expect(t, e).ToBe(nil)
	Expect(t, strings.Trim(out, "\n")).ToBe("2748")
}
func TestClient_Image_by_jpg_from_http(t *testing.T) {
	httpClient := &http.Client{}
	imgURL := "http://images.forbes.com/media/lists/companies/google_200x200.jpg"
	resp, err := httpClient.Get(imgURL)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	client, _ := gosseract.NewClient()
	img, _ := jpeg.Decode(resp.Body)
	out, _ := client.Image(img).Out()
	Expect(t, strings.Trim(out, "\n")).ToBe("Google")
}
func fixImage(fpath string, ext ...string) (img image.Image) {
	f, _ := os.Open(fpath)
	buf, _ := ioutil.ReadFile(f.Name())
	// TODO: refactor
	if len(ext) == 0 {
		img, _ = png.Decode(bytes.NewReader(buf))
	} else {
		img, _ = jpeg.Decode(bytes.NewReader(buf))
	}
	return
}

func TestClient_Out(t *testing.T) {
	client, _ := gosseract.NewClient()
	_, e := client.Out()
	Expect(t, e.Error()).ToBe("Source is not set")
}
