package gosseract

import (
  "github.com/otiai10/gosseract"
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
      Expect(err).To(Equal, nil)
    })
  })
}

func TestServantAllow(t *testing.T) {
  Describe(t, "Allow", func() {
    It("can set whitelist of OCR result chars", func() {
      servant := gosseract.SummonServant()
      servant.Options.Allow(":")
      text, _ := servant.Target("./samples/png/sample002.png").Out()
      Expect(text).To(Equal, "  :  :  \n\n")
    })
  })
}
