package gosseract

import (
	"github.com/otiai10/gosseract"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestHelloGosseract(t *testing.T) {
	Describe(t, "HelloGosseract!!", func() {
		It("should say \"Hello,Gosseract!\"", func() {
			Expect(gosseract.Greeting()).To(Equal, "Hello,Gosseract!")
		})
	})
}

func TestAnyway(t *testing.T) {
	Describe(t, "Anyway", func() {
		It("should exeute OCR anyway with args.", func() {
			args := gosseract.AnywayArgs{
				SourcePath: "samples/png/sample000.png",
			}
			out := gosseract.Anyway(args)
			Expect(out).To(Equal, "01:37:58\n\n")
		})
	})
}
