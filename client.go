package gosseract

import "fmt"

// Client is an client to use gosseract functions
type Client struct {
	tesseract tesseractCmd
	source    path
	result    path
}
type path struct {
	value string
}

func (p *path) Ready() bool {
	return (p.value != "")
}
func (p *path) Get() string {
	return p.value
}

// NewClient provide reference to new Client
func NewClient() (c *Client, e error) {
	tess, e := GetTesseractCmd()
	if e != nil {
		return
	}
	c = &Client{tesseract: tess}
	return
}
func (c *Client) Source(srcPath string) *Client {
	return c
}
func (c *Client) Out() (out string, e error) {
	if e = c.ready(); e != nil {
		return
	}
	out, e = c.execute()
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
func (c *Client) execute() (res string, e error) {
	args := []string{
		c.source.Get(),
	}
	return c.tesseract.Execute(args)
}
