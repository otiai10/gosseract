package gosseract

// #cgo LDFLAGS: -llept -ltesseract
// #include <stdlib.h>
// #include "tessbridge.h"
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

// Version returns the version of Tesseract-OCR
func Version() string {
	api := C.Create()
	defer C.Free(api)
	version := C.Version(api)
	return C.GoString(version)
}

// Client is argument builder for tesseract::TessBaseAPI.
type Client struct {
	api            C.TessBaseAPI
	Trim           bool
	TessdataPrefix *string
	Languages      []string
	ImagePath      string
	Variables      map[string]string
	PageSegMode    *PageSegMode
}

// NewClient construct new Client. It's due to caller to Close this client.
func NewClient() *Client {
	client := &Client{
		api:       C.Create(),
		Variables: map[string]string{},
	}
	return client
}

// Close frees allocated API.
func (c *Client) Close() (err error) {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		err = fmt.Errorf("%v", e)
	// 	}
	// }()
	C.Free(c.api)
	return err
}

// SetImage sets image to execute OCR.
func (c *Client) SetImage(imagepath string) *Client {
	c.ImagePath = imagepath
	return c
}

// SetWhitelist sets whitelist chars.
func (c *Client) SetWhitelist(whitelist string) *Client {
	return c.SetVariable("tessedit_char_whitelist", whitelist)
}

// SetVariable sets parameters.
func (c *Client) SetVariable(key, value string) *Client {
	c.Variables[key] = value
	return c
}

// SetPageSegMode sets PSM
func (c *Client) SetPageSegMode(mode PageSegMode) *Client {
	c.PageSegMode = &mode
	return c
}

// Text finally initalize tesseract::TessBaseAPI, execute OCR and extract text detected as string.
func (c *Client) Text() (string, error) {

	// Defer recover and make error
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	// Initialize tesseract::TessBaseAPI
	if len(c.Languages) == 0 {
		C.Init(c.api, nil, nil)
	} else {
		langs := C.CString(strings.Join(c.Languages, "+"))
		defer C.free(unsafe.Pointer(langs))
		C.Init(c.api, nil, langs)
	}

	// Set Image by giving path
	imagepath := C.CString(c.ImagePath)
	defer C.free(unsafe.Pointer(imagepath))
	C.SetImage(c.api, imagepath)

	for key, value := range c.Variables {
		k, v := C.CString(key), C.CString(value)
		defer C.free(unsafe.Pointer(k))
		defer C.free(unsafe.Pointer(v))
		C.SetVariable(c.api, k, v)
	}

	if c.PageSegMode != nil {
		mode := C.int(*c.PageSegMode)
		C.SetPageSegMode(c.api, mode)
	}

	// Get text by execuitng
	out := C.GoString(C.UTF8Text(c.api))

	// Trim result if needed
	if c.Trim {
		out = strings.Trim(out, "\n")
	}

	return out, err
}
