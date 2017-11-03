package gosseract

import "fmt"

func ExampleNewClient() {
	client := NewClient()
	// Never forget to defer Close. It is due to caller to Close this client.
	defer client.Close()
}

func ExampleClient_SetImage() {
	client := NewClient()
	defer client.Close()

	client.SetImage("./test/data/001-gosseract.png")
	// See "ExampleClient_Text" for more practical usecase ;)
}

func ExampleClient_Text() {

	client := NewClient()
	defer client.Close()

	client.SetImage("./test/data/001-gosseract.png")

	text, err := client.Text()
	fmt.Println(text, err)
	// OUTPUT:
	// otiai10 / gosseract <nil>

}

func ExampleClient_SetWhitelist() {

	client := NewClient()
	defer client.Close()
	client.SetImage("./test/data/002-confusing.png")

	client.SetWhitelist("IO-")
	text1, _ := client.Text()

	client.SetWhitelist("10-")
	text2, _ := client.Text()

	fmt.Println(text1, text2)
	// OUTPUT:
	// IO- IOO 10-100

}
