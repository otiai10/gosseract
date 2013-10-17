package gosseract

import (
  "github.com/otiai10/gosseract-ocr"
  "testing"
)

func TestHelloGosseract(t *testing.T) {
  expected := "Hello,Gosseract!"
  actual := gosseract.HelloGosseract()
  if actual != expected {
    t.Errorf("Expected '%v', Actual '%v'", expected, actual)
  }
}
