![icon](https://f.cloud.github.com/assets/931554/1795347/06dde904-6a07-11e3-8dd3-33b7d466acdb.jpg)

# gosseract-ocr

tesseract-ocr wrapper by Golang

[What is tesseract-ocr?](https://code.google.com/p/tesseract-ocr/) 

# Sample
```go
package main

import (
  "github.com/otiai10/gosseract"
  "fmt"
)

func main() {
  servant := gosseract.SummonServant()

  text, _ := servant.Target("your/image/file.png").Out()

  fmt.Println(text)
}
```

# Set Up
```sh
apt-get install tesseract-ocr # Basic OCR library by C++
# or yum? brew? Choose the way whichever you can install `tesseract-ocr`
go get github.com/otiai10/gosseract
go get github.com/nu7hatch/gouuid
go get github.com/r7kamura/gospel # for testing
```

# First of All, Run the Tests!!

```sh
cd tests
go test -i
go test
```
