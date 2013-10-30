gosseract-ocr
=============

tesseract-ocr wrapper by Golang

[What is tesseract-ocr?](https://code.google.com/p/tesseract-ocr/) 

Set Up
=============

```sh
apt-get install tesseract-ocr # Basic OCR library by C++
# or yum? brew? Choose the way whichever you can install `tesseract-ocr`
go get github.com/otiai10/gosseract-ocr
go get github.com/r7kamura/gosple # for testing
```

At First
=============

```sh
cd tests
go test -i
go test
```

Usage
=============

```go
package main

import (
  "github.com/otiai10/gosseract-ocr"
  "fmt"
)

func main() {
  // sample.png
  text := gosseract.HelloWorld(sample_png)
  fmt.Printf("Result : %v", text)
}
```
