package gosseract

import (
  "github.com/otiai10/gosseract-ocr"
  . "github.com/r7kamura/gospel"
  "testing"
)

func TestHelloServant(t *testing.T) {
  Describe(t, "GosseractServant", func() {
    It("should say \"Hi, I'm gosseract-ocr servant!\"", func() {
      servant := gosseract.SummonServant()
      Expect(servant.Greeting()).To(Equal, "Hi, I'm gosseract-ocr servant!")
    })
  })
}

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
      It("should set Lang and return true(うーん...e == nilみたいにしたいのん...).", func() {
        destination := "eng"// TODO#2: ここengじゃテストにならんでしょうがwww
        Expect(servant.Lang.Use(destination)).To(Equal, true)
        Expect(servant.Lang.Is()).To(Equal, destination)
      })
    })
    Context("with not available language", func() {
      It("should return false(eを返したい).", func() {
        origin := servant.Lang.Is()
        destination := "wrong lang"
        Expect(servant.Lang.Use(destination)).To(Equal, false)
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
        Expect(servant.Options.WithFile(filePath)).To(Equal, true/* TODO#1 */)

        Expect(servant.Options.UseFile).To(Equal, true)
        Expect(servant.Options.FilePath).To(Equal, filePath)
      })
    })
    Context("with existing file", func() {
      It("should not set option file.", func() {
        servant := gosseract.SummonServant()
        filePath := "./not/existing/file/path.txt"

        // Try to Set file
        Expect(servant.Options.WithFile(filePath)).To(Equal, false/* TODO#1 */)

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
        text, err := servant.Eat(filePath).Out()

        Expect(text).To(Equal, "O    \n\n")
        Expect(err).To(Equal, false)
      })

      It("can OCR also without any options.", func() {

        servant := gosseract.SummonServant()
        filePath := "./samples/png/sample000.png"
        text, err := servant.Eat(filePath).Out()

        Expect(text).To(Equal, "01:37:58\n\n")
        Expect(err).To(Equal, false)
      })

    })

  })

}
