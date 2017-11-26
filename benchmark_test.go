package gosseract

import "testing"

func BenchmarkClient_Text(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client := NewClient()
		client.SetImage("./test/data/001-gosseract.png")
		client.Text()
		client.Close()
	}
}
