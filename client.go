package gosseract

// #if __FreeBSD__ >= 10
// #cgo LDFLAGS: -L/usr/local/lib -llept -ltesseract
// #else
// #cgo LDFLAGS: -llept -ltesseract
// #endif
// #include <stdlib.h>
// #include <stdbool.h>
// #include "tessbridge.h"
import "C"
import (
	"fmt"
	"os"
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
	api C.TessBaseAPI

	// Trim specifies characters to trim, which would be trimed from result string.
	// As results of OCR, text often contains unnecessary characters, such as newlines, on the head/foot of string.
	// If `Trim` is set, this client will remove specified characters from the result.
	Trim bool

	// TessdataPrefix can indicate directory path to `tessdata`.
	// It is set `/usr/local/share/tessdata/` or something like that, as default.
	// TODO: Implement and test
	TessdataPrefix *string

	// Languages are languages to be detected. If not specified, it's gonna be "eng".
	Languages []string

	// ImagePath is just path to image file to be processed OCR.
	ImagePath string

	// Variables is just a pool to evaluate "tesseract::TessBaseAPI->SetVariable" in delay.
	// TODO: Think if it should be public, or private property.
	Variables map[string]string

	// PageSegMode is a mode for page layout analysis.
	// See https://github.com/otiai10/gosseract/issues/52 for more information.
	PageSegMode *PageSegMode

	// Config is a file path to the configuration for Tesseract
	// See http://www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version
	// TODO: Fix link to official page
	ConfigFilePath string
}

// NewClient construct new Client. It's due to caller to Close this client.
func NewClient() *Client {
	client := &Client{
		api:       C.Create(),
		Variables: map[string]string{},
		Trim:      true,
	}
	return client
}

// Close frees allocated API. This MUST be called for ANY client constructed by "NewClient" function.
func (client *Client) Close() (err error) {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		err = fmt.Errorf("%v", e)
	// 	}
	// }()
	C.Free(client.api)
	return err
}

// SetImage sets path to image file to be processed OCR.
func (client *Client) SetImage(imagepath string) *Client {
	client.ImagePath = imagepath
	return client
}

// SetLanguage sets languages to use. English as default.
func (client *Client) SetLanguage(langs ...string) *Client {
	client.Languages = langs
	return client
}

// SetWhitelist sets whitelist chars.
// See official documentation for whitelist here https://github.com/tesseract-ocr/tesseract/wiki/ImproveQuality#dictionaries-word-lists-and-patterns
func (client *Client) SetWhitelist(whitelist string) *Client {
	return client.SetVariable("tessedit_char_whitelist", whitelist)
}

// SetVariable sets parameters, representing tesseract::TessBaseAPI->SetVariable.
// See official documentation here https://zdenop.github.io/tesseract-doc/classtesseract_1_1_tess_base_a_p_i.html#a2e09259c558c6d8e0f7e523cbaf5adf5
func (client *Client) SetVariable(key, value string) *Client {
	client.Variables[key] = value
	return client
}

// SetPageSegMode sets "Page Segmentation Mode" (PSM) to detect layout of characters.
// See official documentation for PSM here https://github.com/tesseract-ocr/tesseract/wiki/ImproveQuality#page-segmentation-method
func (client *Client) SetPageSegMode(mode PageSegMode) *Client {
	client.PageSegMode = &mode
	return client
}

// SetConfigFile sets the file path to config file.
func (client *Client) SetConfigFile(fpath string) error {
	info, err := os.Stat(fpath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("the specified config file path seems to be a directory")
	}
	client.ConfigFilePath = fpath
	return nil
}

// It's due to the caller to free this char pointer.
func (client *Client) charLangs() *C.char {
	var langs *C.char
	if len(client.Languages) != 0 {
		langs = C.CString(strings.Join(client.Languages, "+"))
	}
	return langs
}

// It's due to the caller to free this char pointer.
func (client *Client) charConfig() *C.char {
	var config *C.char
	if _, err := os.Stat(client.ConfigFilePath); err == nil {
		config = C.CString(client.ConfigFilePath)
	}
	return config
}

// Initialize tesseract::TessBaseAPI
// TODO: add tessdata prefix
func (client *Client) init() error {
	langs := client.charLangs()
	defer C.free(unsafe.Pointer(langs))
	config := client.charConfig()
	defer C.free(unsafe.Pointer(config))
	res := C.Init(client.api, nil, langs, config)
	if res != 0 {
		// TODO: capture and vacuum stderr from Cgo
		return fmt.Errorf("failed to initialize TessBaseAPI with code %d", res)
	}
	return nil
}

// Prepare tesseract::TessBaseAPI options,
// must be called after `init`.
func (client *Client) prepare() error {
	// Set Image by giving path
	imagepath := C.CString(client.ImagePath)
	defer C.free(unsafe.Pointer(imagepath))
	C.SetImage(client.api, imagepath)

	for key, value := range client.Variables {
		if ok := client.bind(key, value); !ok {
			return fmt.Errorf("failed to set variable with key(%s):value(%s)", key, value)
		}
	}

	if client.PageSegMode != nil {
		mode := C.int(*client.PageSegMode)
		C.SetPageSegMode(client.api, mode)
	}
	return nil
}

// Binds variable to API object.
// Must be called from inside `prepare`.
func (client *Client) bind(key, value string) bool {
	k, v := C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(v))
	res := C.SetVariable(client.api, k, v)
	return bool(res)
}

// Text finally initialize tesseract::TessBaseAPI, execute OCR and extract text detected as string.
func (client *Client) Text() (out string, err error) {
	if err = client.init(); err != nil {
		return
	}
	if err = client.prepare(); err != nil {
		return
	}
	out = C.GoString(C.UTF8Text(client.api))
	if client.Trim {
		out = strings.Trim(out, "\n")
	}
	return out, err
}

// HTML finally initialize tesseract::TessBaseAPI, execute OCR and returns hOCR text.
// See https://en.wikipedia.org/wiki/HOCR for more information of hOCR.
func (client *Client) HOCRText() (out string, err error) {
	if err = client.init(); err != nil {
		return
	}
	if err = client.prepare(); err != nil {
		return
	}
	out = C.GoString(C.HOCRText(client.api))
	return
}
