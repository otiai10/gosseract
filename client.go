package gosseract

// Client is an client to use gosseract functions
type Client struct {
}

// NewClient provide reference to new Client
func NewClient() (*Client, error) {
	return &Client{}, nil
}
