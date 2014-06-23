package gosseract

import "fmt"

// Client is an client to use gosseract functions
type Client struct {
	source path
	result path
}
type path struct {
	value string
}

// NewClient provide reference to new Client
func NewClient() (*Client, error) {
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
	return
}
