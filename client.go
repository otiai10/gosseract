package gosseract

import "os"
import "errors"
import "image"
import "image/png"

// Client of gosseract providing interactive setting
type Client struct {
	source  source
	lang    lang
	options options
}
type source struct {
	FilePath string
	isTmp    bool
	// TODO: accept multiple image formats
}
type lang struct {
	Value      string
	Availables []string
}
type options struct {
	UseFile  bool
	FilePath string
	Digest   map[string]string
}
type VersionInfo struct {
	TesseractVersion string
	GosseractVersion string
}

// Provide new client instance
func NewClient() Client {

	if !tesseractInstalled() {
		panic("Missin `tesseract` command!! install tessearct at first.")
	}

	lang := lang{}
	lang.init()
	opts := options{}
	opts.init()
	return Client{
		lang:    lang,
		options: opts,
	}
}

// Check information of tesseract and gosseract
func (s *Client) Info() VersionInfo {
	tessVersion := getTesseractVersion()
	info := VersionInfo{
		TesseractVersion: tessVersion,
		GosseractVersion: VERSION,
	}
	return info
}

// Give source file to client by file path
func (s *Client) Target(filepath string) *Client {
	// TODO: check existence of this file
	s.source.FilePath = filepath
	return s
}

// Give source file to client by image.Image
func (s *Client) Eat(img image.Image) *Client {
	filepath := genTmpFilePath()
	f, e := os.Create(filepath)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	png.Encode(f, img)

	s.source.FilePath = filepath
	s.source.isTmp = true

	return s
}

// Get result (or error?)
func (s *Client) Out() (string, error) {
	result := execute(s.source.FilePath, s.buildArguments())
	// TODO? : should make `gosseract.client` package?

	if !s.options.UseFile {
		_ = os.Remove(s.options.FilePath)
	}
	if s.source.isTmp {
		_ = os.Remove(s.source.FilePath)
	}

	// TODO: handle errors
	return result, nil
}

// Make up arguments appropriate to tesseract command
func (s *Client) buildArguments() []string {
	var args []string
	args = append(args, "-l", s.lang.Value)
	if !s.options.UseFile {
		s.options.FilePath = makeUpOptionFile(s.options.Digest)
	}
	args = append(args, s.options.FilePath)
	return args
}

// Make up option file for tesseract command.
// (is needless if tesseract accepts such options by cli options)
func makeUpOptionFile(digestMap map[string]string) (fpath string) {
	fpath = ""
	var digestFileContents string
	for k, v := range digestMap {
		digestFileContents = digestFileContents + k + " " + v + "\n"
	}
	if digestFileContents == "" {
		return fpath
	}
	fpath = genTmpFilePath()
	f, _ := os.Create(fpath)
	defer f.Close()
	_, _ = f.WriteString(digestFileContents)
	return fpath
}

func (s *Client) LangAvailable() []string {
	return s.lang.Availables
}
func (s *Client) LangHave(key string) bool {
	for _, language := range s.lang.Availables {
		if language == key {
			return true
		}
	}
	return false
}
func (s *Client) LangIs() string {
	return s.lang.Value
}
func (s *Client) LangUse(key string) error {
	if s.LangHave(key) {
		s.lang.Value = key
		return nil
	}
	return errors.New("Language `" + key + "` is not available.")
}

func (l *lang) init() *lang {
	l.Value = "eng" // "eng" in default
	l.Availables = getAvailables()
	return l
}

func getAvailables() []string {
	langs := []string{}
	for _, lang := range getAvailableLanguages() {
		langs = append(langs, lang)
	}
	return langs
}

func (o *options) init() *options {
	o.UseFile = false
	o.FilePath = ""
	o.Digest = make(map[string]string)
	return o
}

func (s *Client) OptionWithFile(path string) error {
	_, e := os.Open(path)
	if e != nil {
		return errors.New("No such option file `" + path + "` is found.")
	}
	s.options.UseFile = true
	s.options.FilePath = path
	return nil
}

func (s *Client) AllowChars(charAllowed string) {
	if charAllowed == "" {
		return
	}
	s.options.Digest["tessedit_char_whitelist"] = charAllowed
	return
}
