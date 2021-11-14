# gosseract OCR

[![Go Test](https://github.com/otiai10/gosseract/actions/workflows/go-ci.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/go-ci.yml)
[![Docker Test](https://github.com/otiai10/gosseract/actions/workflows/runtime-docker.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/runtime-docker.yml)
[![Vagrant Test](https://github.com/otiai10/gosseract/actions/workflows/runtime-vagrant.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/runtime-vagrant.yml)
[![codecov](https://codecov.io/gh/otiai10/gosseract/branch/main/graph/badge.svg)](https://codecov.io/gh/otiai10/gosseract)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/gosseract)](https://goreportcard.com/report/github.com/otiai10/gosseract)
[![Maintainability](https://api.codeclimate.com/v1/badges/351d9027a3c517505094/maintainability)](https://codeclimate.com/github/otiai10/gosseract/maintainability)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/otiai10/gosseract/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/otiai10/gosseract/v2.svg)](https://pkg.go.dev/github.com/otiai10/gosseract/v2)

Golang OCR package, by using Tesseract C++ library.

# OCR Server

Do you just want OCR server, or see the working example of this package? Yes, there is already-made server application, which is seriously easy to deploy!

👉 https://github.com/otiai10/ocrserver

# Example

```go
package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("path/to/image.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
```

# Install

1. [tesseract-ocr](https://github.com/tesseract-ocr/tessdoc), including library and headers
2. `go get -t github.com/otiai10/gosseract/v2`

Check [Dockerfile](https://github.com/otiai10/gosseract/blob/main/Dockerfile) for more detail of installation, or you can just try by `docker run -it --rm otiai10/gosseract`.

# Test

In case you have [tesseract-ocr](https://github.com/tesseract-ocr/tessdoc) on your local, you can just hit

```
% go test .
```

Otherwise, if you **DON'T** want to install tesseract-ocr on your local, kick `./test/runtime` which is using Docker and Vagrant to test the source code on some runtimes.

```
% ./test/runtime --driver docker
% ./test/runtime --driver vagrant
```

Check [./test/runtimes](https://github.com/otiai10/gosseract/tree/main/test/runtimes) for more information about runtime tests.

# Issues

- [https://github.com/otiai10/gosseract/issues](https://github.com/otiai10/gosseract/issues?utf8=%E2%9C%93&q=is%3Aissue)
