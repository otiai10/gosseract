# gosseract

Golang OCR package, wrapping Tesseract-OCR C++ library.


# OCR Server

Do you just want OCR server, or see the working example of this package? Yes, there is already-made server application, which is seriously easy to deploy!

👉 https://github.com/otiai10/ocrserver

# Example

```go
package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
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

1. [tesseract](https://github.com/tesseract-ocr/tesseract/wiki), including library and headers
2. `go get github.com/otiai10/gosseract`

# Test

For basic test, install [mint](https://github.com/otiai10/mint) by `go get github.com/otiai10/mint` then `go test`. It requires tesseract-ocr and its library and header files installed on local machine.

```
% go get -u github.com/otiai10/mint
% go test`.
```

If you don't want to install tesseract-ocr on your local machine, use `./test/script/runtime.sh` and use Docker runtime (and Vagrant coming soon) to test the source code.

```
% ./test/script/runtime.sh --driver docker
```
