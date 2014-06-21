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
	servant := gosseract.SummonServant()

	text, _ := servant.Target("your/image/file.png").Out()

	fmt.Println(text)
}

```

# setup
```sh
apt-get install tesseract-ocr # Basic OCR library by C++
# or yum? brew? Choose the way whichever you can install `tesseract-ocr`
# for Mac OS X 'brew install tesseract'
go get github.com/otiai10/gosseract
```

# test
```sh
go test ./...
```

# issues
- https://github.com/otiai10/gosseract/issues?state=open
