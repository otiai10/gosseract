package gosseract

import "fmt"
import "os/exec"
import "bytes"
import "regexp"

// Client is an client to use gosseract functions
type Client struct {
	source path
	result path
}
type path struct {
	value string
}

func (p *path) Ready() bool {
	return (p.value != "")
}

// NewClient provide reference to new Client
func NewClient() (*Client, error) {
	v, e := Version()
	fmt.Println(v, e)
	return &Client{}, nil
}
func (c *Client) Source(srcPath string) *Client {
	return c
}
func (c *Client) Out() (out string, e error) {
	if e = c.ready(); e != nil {
		return
	}
	out = "gosseract"
	return
}
func (c *Client) Must(params map[string]string) (out string, e error) {
	if e = c.accept(params); e != nil {
		return
	}
	return c.Out()
}
func (c *Client) accept(params map[string]string) (e error) {
	var ok bool
	var src string
	if src, ok = params["src"]; !ok {
		return fmt.Errorf("Missing parameter `src`.")
	}
	c.source = path{src}
	return
}
func (c *Client) ready() (e error) {
	if !c.source.Ready() {
		return fmt.Errorf("Source is not set")
	}
	return
}
func Version() (v string, e error) {
	v, e = execTesseractCommandWithStderr("--version")
	if e != nil {
		return
	}
	exp := regexp.MustCompile("^tesseract ([0-9\\.]+)")
	matches := exp.FindStringSubmatch(v)
	if len(matches) < 2 {
		e = fmt.Errorf("tesseract version not found: response is `%s`.", v)
	}
	v = matches[1]
	return
}
func execTesseractCommandWithStderr(opt string) (res string, e error) {
	cmd := exec.Command("tesseract", opt)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if e = cmd.Run(); e != nil {
		return
	}
	res = stderr.String()
	return
}
