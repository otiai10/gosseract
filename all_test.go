package gosseract_test

import "github.com/otiai10/gosseract"
import . "github.com/otiai10/mint"
import "testing"

func Test_Greet(t *testing.T) {
	Expect(t, gosseract.Greet()).ToBe("Hello,Gosseract.")
}

func Test_Must(t *testing.T) {
	params := map[string]string{
		"src": "./samples/hoge.png",
	}
	Expect(t, gosseract.Must(params)).ToBe("gosseract")
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
