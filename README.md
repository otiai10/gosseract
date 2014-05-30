This project is under reconstructing in tree [`wip/v2`](https://github.com/otiai10/gosseract/tree/wip/v2).

- More effective error handling, without using `panic`.
- More plain implementations, to make it easy to read.
- etc...

[![Build Status](https://travis-ci.org/otiai10/gosseract.svg?branch=develop)](https://travis-ci.org/otiai10/gosseract)

# gosseract-ocr

[Tesseract-OCR](https://code.google.com/p/tesseract-ocr/) command wrapper for Golang

# example
```go
package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	servant := gosseract.SummonServant()

	text, _ := servant.Target("your/image/file.png").Out()

	fmt.Println(text)
}
```

# setup
```sh
apt-get install tesseract-ocr # Basic OCR library by C++
# or yum? brew? Choose the way whichever you can install `tesseract-ocr`
go get github.com/otiai10/gosseract
```

# test
```sh
go test ./...
```

# issues
- https://github.com/otiai10/gosseract/issues?state=open
