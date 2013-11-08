package gosseract

import (
  "github.com/otiai10/gosseract"
  "fmt"

  "image"
)

func ExampleAnyway() {
  args := gosseract.AnywayArgs{
    SourcePath: "samples/png/sample000.png",
  }
  out := gosseract.Anyway(args)
  fmt.Println(out)
}

func ExampleTarget() {
  var sourceFilePath string
  sourceFilePath = "./samples/png/sample000.png"

  servant := gosseract.SummonServant()
  text, err := servant.Target(sourceFilePath).Out()

  if err != nil {
    panic(err)
  }
  fmt.Println(text)
}

func ExampleOptionWithFile() {
  var optionFilePath string
  optionFilePath = "./samples/option/digest001.txt"

  var sourceFilePath string
  sourceFilePath = "./samples/png/sample000.png"

  servant := gosseract.SummonServant()
  servant.OptionWithFile(optionFilePath)
  text, err := servant.Target(sourceFilePath).Out()

  if err != nil {
    panic(err)
  }
  fmt.Println(text)
}

func ExampleEat() {
  var img image.Image
  img = fixtureImageObj("./samples/png/sample001.png")

  servant := gosseract.SummonServant()
  text, err := servant.Eat(img).Out()

  if err != nil {
    panic(err)
  }
  fmt.Println(text)
}
