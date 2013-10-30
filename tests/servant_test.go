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

  Describe(t, "AvailableLanguages", func() {
    It("should give available languages of Tesseract.", func() {
      servant := gosseract.NewServant()
      langs := servant.AvailableLanguages()
      Expect(len(langs)).To(NotEqual, 0)
    })
    It("should contain 'eng' at least.", func() {
      servant := gosseract.NewServant()
      langs := servant.AvailableLanguages()
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
}
