# gosseract-ocr

[Tesseract-OCR](https://code.google.com/p/tesseract-ocr/) command wrapper for Golang

[![Build Status](https://travis-ci.org/otiai10/gosseract.svg?branch=develop)](https://travis-ci.org/otiai10/gosseract)

# example
```go
package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
    // get client
	client, _ := gosseract.NewClient()
    // pass path to source image
	text, _ := client.Target("your/image/file.png").Out()

	fmt.Println(text)
}

```

# dependencies

- [tesseract-ocr](https://code.google.com/p/tesseract-ocr/)#3.02~

# test
```sh
go test ./...
```

# issues
- https://github.com/otiai10/gosseract/issues?state=open
