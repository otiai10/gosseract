# gosseract OCR

[![Go Test](https://github.com/otiai10/gosseract/actions/workflows/go-ci.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/go-ci.yml)
[![Docker Test](https://github.com/otiai10/gosseract/actions/workflows/runtime-docker.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/runtime-docker.yml)
[![BSD Test](https://github.com/otiai10/gosseract/actions/workflows/runtime-vmactions.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/runtime-vmactions.yml)
[![Windows Test](https://github.com/otiai10/gosseract/actions/workflows/windows-ci.yml/badge.svg)](https://github.com/otiai10/gosseract/actions/workflows/windows-ci.yml)
[![codecov](https://codecov.io/gh/otiai10/gosseract/branch/main/graph/badge.svg)](https://codecov.io/gh/otiai10/gosseract)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/gosseract)](https://goreportcard.com/report/github.com/otiai10/gosseract)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/otiai10/gosseract/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/otiai10/gosseract/v2.svg)](https://pkg.go.dev/github.com/otiai10/gosseract/v2)

Golang OCR package, by using [Tesseract](https://github.com/tesseract-ocr/tesseract) C++ library.

# OCR Server

If you need an [OCR server](https://github.com/otiai10/ocrserver) or want to see a working example of this package, there is a ready-made server application, which is very easy to deploy!

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

# Installation

## macOS

```bash
brew install tesseract
go get -t github.com/otiai10/gosseract/v2
```

## Linux (Debian/Ubuntu)

```bash
sudo apt-get install -y libtesseract-dev libleptonica-dev tesseract-ocr-eng
go get -t github.com/otiai10/gosseract/v2
```

Please check this [Dockerfile](https://github.com/otiai10/gosseract/blob/main/Dockerfile) to get started.
Alternatively, you can deploy the pre-existing Docker image by invoking `docker run -it --rm otiai10/gosseract`.

## Windows

Windows support requires [vcpkg](https://vcpkg.io/) and [MinGW-w64](https://www.mingw-w64.org/):

```bash
# Install Tesseract via vcpkg
vcpkg install tesseract:x64-windows

# Create MinGW import libraries from vcpkg DLLs
cd C:/vcpkg/installed/x64-windows/bin
gendef tesseract55.dll leptonica-1.87.0.dll
dlltool -d tesseract55.def -l libtesseract.a -D tesseract55.dll
dlltool -d leptonica-1.87.0.def -l libleptonica.a -D leptonica-1.87.0.dll
mv *.a ../lib/

# Download language data
mkdir C:/tessdata
curl -L -o C:/tessdata/eng.traineddata https://github.com/tesseract-ocr/tessdata/raw/main/eng.traineddata
```

Set environment variables before building:

```bash
export CGO_ENABLED=1
export CC=C:/mingw64/bin/gcc.exe
export CGO_CFLAGS="-IC:/vcpkg/installed/x64-windows/include"
export CGO_LDFLAGS="-LC:/vcpkg/installed/x64-windows/lib"
export TESSDATA_PREFIX="C:/tessdata"
export PATH="/c/mingw64/bin:/c/vcpkg/installed/x64-windows/bin:$PATH"
```

For detailed troubleshooting, see [knowledge/windows-support.md](./knowledge/windows-support.md).

# Test

In case you have [tesseract-ocr](https://github.com/tesseract-ocr/tessdoc) installed on your local environment, you can run the tests with:

```
% go test .
```

If you **DON'T** want to install tesseract-ocr on your local environment, run `./test/runtime` which utilises Docker and Vagrant to test the source code on some runtimes.

```
% ./test/runtime --engine docker
% ./test/runtime --engine vagrant
```

Check [./test/runtimes](https://github.com/otiai10/gosseract/tree/main/test/runtimes) for more information about runtime tests.

> **Note**: Clear Linux support was removed in January 2026 as [Intel discontinued the distribution in July 2025](https://www.phoronix.com/news/Intel-Linux-News-2025).

> **Note**: Arch Linux support was removed in January 2026 as the official image does not provide ARM64 architecture support.

# Issues

- [https://github.com/otiai10/gosseract/issues](https://github.com/otiai10/gosseract/issues?utf8=%E2%9C%93&q=is%3Aissue)
