package gosseract

import (
  "github.com/otiai10/gosseract"
  . "github.com/r7kamura/gospel"
  "testing"

  "image"
  "image/png"
  "os"
  "io/ioutil"
  "bytes"

  "reflect"
)

func TestServant(t *testing.T) {
  Describe(t, "Info", func() {
    It("shoul show version of Tesseract and Gosseract.", func() {
      servant := gosseract.SummonServant()
      info := servant.Info()
      Expect(info.GosseractVersion).To(Equal, "0.0.1")
      Expect(info.TesseractVersion).To(Exist)
    })
  })
}

func TestServantLang(t *testing.T) {
  Describe(t, "Available", func() {
    It("should give available languages of Tesseract", func() {
      servant := gosseract.SummonServant()
      langs := servant.Lang.Available()
      Expect(len(langs)).To(NotEqual, 0)
    })
    It("should contain 'eng' at least.", func() {
      servant := gosseract.SummonServant()
      langs := servant.Lang.Available()
      containEng := false
      for _,lang := range langs {
        if lang == "eng" {
          containEng = true
          break
        }
      }
      Expect(containEng).To(Equal, true)
    })
  })
  Describe(t, "Have", func() {
    It("should give whether argument language is available or not.", func() {
      servant := gosseract.SummonServant()
      Expect(servant.Lang.Have("eng")).To(Equal, true)
    })
  })
  Describe(t, "Is", func() {
    It("should give current language setting.", func() {
      servant := gosseract.SummonServant()
      Expect(servant.Lang.Is()).To(Equal, "eng")
    })
  })
  Describe(t, "Use", func() {
    servant := gosseract.SummonServant()
    Context("with available language", func() {
      It("should set Lang and return nil.", func() {
        destination := "eng"// TODO#2: ここengじゃテストにならんでしょうがwww
        Expect(servant.Lang.Use(destination)).To(Equal, nil)
        Expect(servant.Lang.Is()).To(Equal, destination)
      })
    })
    Context("with not available language", func() {
      It("should return error.", func() {
        origin := servant.Lang.Is()
        destination := "wrong lang"
        e := servant.Lang.Use(destination)
        Expect(reflect.TypeOf(e).String()).To(Equal, "*errors.errorString")
        Expect(servant.Lang.Is()).To(NotEqual, destination)
        Expect(servant.Lang.Is()).To(Equal, origin)
      })
    })
  })
}

func TestServantOptions(t *testing.T) {
  Describe(t, "WithFile", func() {
    Context("with existing file", func() {
      It("should set option file.", func() {
        servant := gosseract.SummonServant()
        // Do not use file in default
        Expect(servant.Options.FilePath).To(Equal, "")
        Expect(servant.Options.UseFile).To(Equal, false)
        filePath := "./samples/option/digest000.txt"

        // Try to Set file
        Expect(servant.Options.WithFile(filePath)).To(Equal, nil)

        Expect(servant.Options.UseFile).To(Equal, true)
        Expect(servant.Options.FilePath).To(Equal, filePath)
      })
    })
    Context("with existing file", func() {
      It("should not set option file.", func() {
        servant := gosseract.SummonServant()
        filePath := "./not/existing/file/path.txt"

        // Try to Set file
        e := servant.Options.WithFile(filePath)
        Expect(reflect.TypeOf(e).String()).To(Equal, "*errors.errorString")

        Expect(servant.Options.FilePath).To(Equal, "")
        Expect(servant.Options.UseFile).To(Equal, false)
      })
    })
  })
}

func TestServantStory(t *testing.T) {
  Describe(t, "Usage of Servant, Servant", func() {

    Context("with option file", func() {

      It("can OCR according to option file.", func() {

        servant := gosseract.SummonServant()
        servant.Options.WithFile("./samples/option/digest001.txt")
        filePath := "./samples/png/sample000.png"
        text, err := servant.Target(filePath).Out()

        Expect(text).To(Equal, "O    \n\n")
        Expect(err).To(Equal, nil)
      })

      It("can OCR also without any options.", func() {

        servant := gosseract.SummonServant()
        filePath := "./samples/png/sample000.png"
        text, err := servant.Target(filePath).Out()

        Expect(text).To(Equal, "01:37:58\n\n")
        Expect(err).To(Equal, nil)
      })

    })

  })

}

func fixtureImageObj(fpath string) image.Image {
  f, _ := os.Open(fpath)
  buf, _ := ioutil.ReadFile(f.Name())
  img, _ := png.Decode(bytes.NewReader(buf))
  // *image.RGBA
  return img
}

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
