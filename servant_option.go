package gosseract

import (
	"errors"
	"os"
)

func (o *options) init() *options {
	o.UseFile = false
	o.FilePath = ""
	o.Digest = make(map[string]string)
	return o
}

func (s *Servant) OptionWithFile(path string) error {
	_, e := os.Open(path)
	if e != nil {
		return errors.New("No such option file `" + path + "` is found.")
	}
	s.options.UseFile = true
	s.options.FilePath = path
	return nil
}

func (s *Servant) AllowChars(charAllowed string) {
	if charAllowed == "" {
		return
	}
	s.options.Digest["tessedit_char_whitelist"] = charAllowed
	return
}
