package gosseract_test

import "github.com/otiai10/gosseract"
import "testing"
import "fmt"

func ExampleMust(t *testing.T) {
	out := gosseract.Must("./.samples/png/sample002.png")
	fmt.Println(out)
}
