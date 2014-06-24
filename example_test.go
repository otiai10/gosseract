package gosseract_test

import "github.com/otiai10/gosseract"
import "testing"
import "fmt"

func ExampleMust(t *testing.T) {
	// TODO: it panics! error handling in *Client.accept
	out := gosseract.Must(map[string]string{"src": "./.samples/png/sample002.png"})
	fmt.Println(out)
}
