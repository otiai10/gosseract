gosseract-ocr
=============

Tesseract-ocr wrapper by Golang

To set up
=============

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
  /*
   * sample.png
   */
  text := gosseract.HelloWorld(sample_png)
  fmt.Printf("Result : %v", text)
}
```
