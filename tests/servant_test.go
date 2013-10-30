package gosseract

import (
  "github.com/otiai10/gosseract-ocr"
  . "github.com/r7kamura/gospel"
  "testing"
)

func TestHelloServant(t *testing.T) {
  Describe(t, "GosseractServant", func() {
    It("should say \"Hi, I'm gosseract-ocr servant!\"", func() {
      servant := gosseract.NewServant()
      Expect(servant.Greeting()).To(Equal, "Hi, I'm gosseract-ocr servant!")
    })
  })
}

func TestServant(t *testing.T) {
  Describe(t, "Info", func() {
    It("shoul show version of Tesseract and Gosseract.", func() {
      servant := gosseract.NewServant()
      info := servant.Info()
      Expect(info.GosseractVersion).To(Equal, "0.0.1")
      Expect(info.TesseractVersion).To(Exist)
    })
  })
}
