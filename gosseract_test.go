package gosseract_test

import "github.com/otiai10/gosseract"
import "testing"
import "fmt"
import "os"

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if expected != actual {
		fmt.Printf("`%+v` expected, but `%+v` actual.\n", expected, actual)
		t.Fail()
		os.Exit(1)
	}
}

func TestGosseract_Greeting(t *testing.T) {
    assert(
        t,
        gosseract.Greeting(),
        "Hello,Gosseract!",
    )
}

func TestGosseract_Anyway(t *testing.T) {
    args := gosseract.AnywayArgs{
        SourcePath: ".samples/png/sample000.png",
    }
    assert(
        t,
        gosseract.Anyway(args),
        "01:37:58\n\n",
    )
}
