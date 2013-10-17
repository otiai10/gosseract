gosseract-ocr
=============

tesseract-ocr wrapper by Golang

[What is tesseract-ocr?](https://code.google.com/p/tesseract-ocr/) 

Set Up

```
% apt-get install tesseract-ocr
% go get github.com/otiai10/gosseract-ocr
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
