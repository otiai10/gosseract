package gosseract

import (
  "github.com/otiai10/gosseract"
  . "github.com/r7kamura/gospel"
  "testing"

  "image"
)

func ExampleAnyway(t *testing.T) {
  Describe(t, "gosseract.Anyway", func() {
    It("should exeute OCR anyway with args.", func() {
      args := gosseract.AnywayArgs{
        SourcePath: "samples/png/sample000.png",
      }
      out := gosseract.Anyway(args)
      Expect(out).To(Equal, "01:37:58\n\n")
    })
  })
}

func ExampleTarget(t *testing.T) {
  Describe(t, "Servant.Target", func() {
    It("can OCR also without any options.", func() {

      var sourceFilePath string
      sourceFilePath = "./samples/png/sample000.png"

      servant := gosseract.SummonServant()
      text, err := servant.Target(sourceFilePath).Out()

      Expect(text).To(Equal, "01:37:58\n\n")
      Expect(err).To(Equal, nil)
    })
  })
}

func ExampleOptionWithFile(t *testing.T) {
  Describe(t, "Servant.Target", func() {
    Context("with option file", func() {
      It("can OCR according to option file.", func() {

        var optionFilePath string
        optionFilePath = "./samples/option/digest001.txt"

        var sourceFilePath string
        sourceFilePath = "./samples/png/sample000.png"

        servant := gosseract.SummonServant()
        servant.OptionWithFile(optionFilePath)
        text, err := servant.Target(sourceFilePath).Out()

        Expect(text).To(Equal, "O    \n\n")
        Expect(err).To(Equal, nil)
      })
    })
  })
}

func ExampleEat(t *testing.T) {
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
