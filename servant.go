package gosseract

type Servant struct {
}

func NewServant() Servant {
  return Servant{}
}

func (s *Servant) Greeting() string {
  return "Hi, I'm gosseract-ocr servant!"
}
