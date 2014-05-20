package gosseract_test

import "github.com/otiai10/gosseract"
import "testing"
import "fmt"
import "os"

func notNil(t *testing.T, actual interface{}) {
    if actual == nil {
		fmt.Printf("`%+v` expected to be not nil", actual)
		t.Fail()
		os.Exit(1)
    }
}

func TestServant_Info(t *testing.T) {
    servant := gosseract.SummonServant()
    info := servant.Info()
    assert(t, info.GosseractVersion, "0.0.1")
    notNil(t, info.TesseractVersion)
}
