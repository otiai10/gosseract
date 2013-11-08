package gosseract

import (
  "os"
  "image"
  "image/png"
)

// Servant of gosseract providing interactive setting
type Servant struct {
  source  source
  lang    lang
  options options
}
type source struct {
  FilePath string
  // 画像形式とかくるんだろうな今後
}
type lang struct {
  Value      string
  Availables []string
}
type options struct {
  UseFile   bool
  FilePath  string
  Digest map[string]string
}
type VersionInfo struct {
  TesseractVersion string
  GosseractVersion string
}

// Provide new servant instance
func SummonServant() Servant {

  if ! tesseractInstalled() {
    panic("Missin `tesseract` command!! install tessearct at first.")
  }

  lang := lang{}
  lang.init()
  opts := options{}
  opts.init()
  return Servant{
    lang:    lang,
    options: opts,
  }
}

// Check information of tesseract and gosseract
func (s *Servant) Info() VersionInfo {
  tessVersion := getTesseractVersion()
  info := VersionInfo{
    TesseractVersion: tessVersion,
    GosseractVersion: VERSION,
  }
  return info
}

// Give source file to servant by file path
func (s *Servant) Target(filepath string) *Servant {
  // TODO: ファイル存在チェック
  s.source.FilePath = filepath
  return s
}

// Give source file to servant by image.Image
func (s *Servant) Eat(img image.Image) *Servant {
  filepath := genTmpFilePath()
  f, e := os.Create(filepath)
  if e != nil {
    panic(e)
  }
  defer f.Close()
  png.Encode(f, img)

  s.source.FilePath = filepath

  return s
}

// Get result (or error?)
func (s *Servant) Out() (string, error) {
  result := execute(s.source.FilePath, s.buildArguments())
  // errorここハードにnilなら要らなくないか？
  return result, nil
}

func (s *Servant) buildArguments() []string {
  var args []string
  args = append(args, "-l", s.lang.Value)
  if ! s.options.UseFile {
    s.options.FilePath = makeUpOptionFile(s.options.Digest)
  }
  args = append(args, s.options.FilePath)
  return args
}
func makeUpOptionFile(digestMap map[string]string) (fpath string) {
  fpath = ""
  var digestFileContents string
  for k, v := range digestMap  {
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
