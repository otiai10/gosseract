package gosseract

import (
  "github.com/otiai10/gosseract-ocr"
  . "github.com/r7kamura/gospel"
  "testing"
  "image"
)

func TestServantEat(t *testing.T) {
  Describe(t, "Eat", func() {
    It("can OCR from `image.Image`.", func() {

      var img image.Image
      img = fixtureImageObj("./samples/png/sample001.png")

      servant := gosseract.SummonServant()
      text, err := servant.Eat(img).Out()
      Expect(text).To(Equal, "03:41:26\n\n")
      Expect(err).To(Equal, false)
    })
  })
}
