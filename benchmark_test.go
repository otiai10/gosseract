package gosseract

import "testing"

func BenchmarkClient_Text(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client := NewClient()
		client.SetImage("./test/data/001-helloworld.png")
		client.Text()
		client.Close()
	}
}

func BenchmarkClient_Text2(b *testing.B) {
	client := NewClient()
	for i := 0; i < b.N; i++ {
		client.SetImage("./test/data/001-helloworld.png")
		client.Text()
	}
	client.Close()
}

func BenchmarkClient_GetBoundingBoxes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client := NewClient()
		client.SetImage("./test/data/003-longer-text.png")
		client.GetBoundingBoxes(3)
		client.Close()
	}
}

func BenchmarkClient_GetBoundingBoxesVerbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client := NewClient()
		client.SetImage("./test/data/003-longer-text.png")
		client.GetBoundingBoxesVerbose()
		client.Close()
	}
}
