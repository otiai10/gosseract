# Gosseract-OCR [![Build Status](https://travis-ci.org/otiai10/gosseract.svg?branch=master)](https://travis-ci.org/otiai10/gosseract) [![GoDoc](https://godoc.org/github.com/otiai10/gosseract?status.png)](https://godoc.org/github.com/otiai10/gosseract)

[Tesseract-OCR](https://github.com/tesseract-ocr/tesseract) command wrapper for Golang

# code example

```go
package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
    // This is the simlest way :)
    out := gosseract.Must(gosseract.Params{
			Src:       "your/img/file.png",
			Languages: "eng+heb",
    })
    fmt.Println(out)

    // Using client
    client, _ := gosseract.NewClient()
    out, _ = client.Src("your/img/file.png").Out()
    fmt.Println(out)
}
```

# sample application

[![ocrserver](https://github.com/otiai10/ocrserver/raw/master/assets/favicon.png)](https://github.com/otiai10/ocrserver)
[ocrserver](https://github.com/otiai10/ocrserver): the minimum OCR server with using gosseract.


# installation

1. install [tesseract-ocr](https://github.com/tesseract-ocr/tesseract)
2. install [go](http://golang.org/doc/install)
3. install [gosseract](https://godoc.org/github.com/otiai10/gosseract)
    - `go get github.com/otiai10/gosseract`
4. install [mint for testing](https://godoc.org/github.com/otiai10/mint)
    - `go get github.com/otiai10/mint`
5. run the tests at first↓

# test
```sh
go test ./...
```

# dependencies

- [tesseract-ocr](https://github.com/tesseract-ocr/tesseract)#3.02~
- [mint](https://github.com/otiai10/mint) to simplize tests

# issues
- https://github.com/otiai10/gosseract/issues?state=open
