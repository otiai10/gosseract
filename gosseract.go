package gosseract

// Greet is a test func :)
func Greet() string {
	return "Hello,Gosseract."
}

func Must(params map[string]string) (out string) {
	client, _ := NewClient()
	out, _ = client.Must(params)
	return
}
