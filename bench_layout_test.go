package gosseract_test

import (
	"fmt"
	"image"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

// ---------- Stand-ins for CGO types so this compiles without CGO ----------
type TessBaseAPI = unsafe.Pointer
type PixImage = unsafe.Pointer

// ---------- ORIGINAL STRUCTS ----------
type ClientOrig struct {
	api            TessBaseAPI
	pixImage       PixImage
	Trim           bool
	TessdataPrefix string
	Languages      []string
	Variables      map[string]string // simplified: SettableVariable -> string
	ConfigFilePath string
	shouldInit     bool
}

type BoundingBoxOrig struct {
	Box        image.Rectangle
	Word       string
	Confidence float64
	BlockNum   int
	ParNum     int
	LineNum    int
	WordNum    int
}

// ---------- OPTIMIZED (field-reordered) ----------
type ClientOpt struct {
	api            TessBaseAPI
	pixImage       PixImage
	TessdataPrefix string
	Languages      []string
	Variables      map[string]string
	ConfigFilePath string

	Trim       bool
	shouldInit bool
}

// ---------- SLIM (right-sized numeric types, better alignment) ----------
type BoundingBoxSlim struct {
	Word       string
	X0, Y0     int32
	X1, Y1     int32
	Confidence float32
	BlockNum   int32
	ParNum     int32
	LineNum    int32
	WordNum    int32
}

func (b BoundingBoxSlim) Rect() image.Rectangle {
	return image.Rect(int(b.X0), int(b.Y0), int(b.X1), int(b.Y1))
}

func FromRectangle(word string, r image.Rectangle, conf float32, block, par, line, wordNum int) BoundingBoxSlim {
	return BoundingBoxSlim{
		Word:       word,
		X0:         int32(r.Min.X),
		Y0:         int32(r.Min.Y),
		X1:         int32(r.Max.X),
		Y1:         int32(r.Max.Y),
		Confidence: conf,
		BlockNum:   int32(block),
		ParNum:     int32(par),
		LineNum:    int32(line),
		WordNum:    int32(wordNum),
	}
}

// ---------- Helpers ----------
func randWord(rng *rand.Rand) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	n := 3 + rng.Intn(9)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}

func makeBBoxOrig(n int) []BoundingBoxOrig {
	rng := rand.New(rand.NewSource(42))
	out := make([]BoundingBoxOrig, n)
	for i := range out {
		out[i].Box = image.Rect(rng.Intn(1000), rng.Intn(1000), rng.Intn(1000)+10, rng.Intn(1000)+10)
		out[i].Word = randWord(rng)
		out[i].Confidence = rng.Float64()
		out[i].BlockNum = i
		out[i].ParNum = i % 7
		out[i].LineNum = i % 11
		out[i].WordNum = i % 13
	}
	return out
}

func makeBBoxSlim(n int) []BoundingBoxSlim {
	rng := rand.New(rand.NewSource(42))
	out := make([]BoundingBoxSlim, n)
	for i := range out {
		x0 := rng.Int31n(1000)
		y0 := rng.Int31n(1000)
		out[i].X0, out[i].Y0 = x0, y0
		out[i].X1, out[i].Y1 = x0+10, y0+10
		out[i].Word = randWord(rng)
		out[i].Confidence = rng.Float32()
		out[i].BlockNum = int32(i)
		out[i].ParNum = int32(i % 7)
		out[i].LineNum = int32(i % 11)
		out[i].WordNum = int32(i % 13)
	}
	return out
}

func makeClientOrig(n int) []ClientOrig {
	out := make([]ClientOrig, n)
	for i := range out {
		out[i].TessdataPrefix = "/usr/share/tessdata"
		out[i].Languages = []string{"eng", "osd"}
		out[i].Variables = map[string]string{"tessedit_char_whitelist": "abc123"}
		out[i].ConfigFilePath = "tesseract.conf"
		out[i].Trim = (i%2 == 0)
		out[i].shouldInit = (i%3 == 0)
	}
	return out
}

func makeClientOpt(n int) []ClientOpt {
	out := make([]ClientOpt, n)
	for i := range out {
		out[i].TessdataPrefix = "/usr/share/tessdata"
		out[i].Languages = []string{"eng", "osd"}
		out[i].Variables = map[string]string{"tessedit_char_whitelist": "abc123"}
		out[i].ConfigFilePath = "tesseract.conf"
		out[i].Trim = (i%2 == 0)
		out[i].shouldInit = (i%3 == 0)
	}
	return out
}

// Prevents compiler from optimizing away work.
var blackhole any

// ---------- Size reporting ----------
func TestSizes(t *testing.T) {
	t.Logf("sizeof(ClientOrig)  = %d", unsafe.Sizeof(ClientOrig{}))
	t.Logf("sizeof(ClientOpt)   = %d", unsafe.Sizeof(ClientOpt{}))
	t.Logf("sizeof(BBoxOrig)    = %d", unsafe.Sizeof(BoundingBoxOrig{}))
	t.Logf("sizeof(BBoxSlim)    = %d", unsafe.Sizeof(BoundingBoxSlim{}))
}

// ---------- Heap footprint tests ----------
func TestHeapFootprint(t *testing.T) {
	const N = 2_000_000
	var ms1, ms2 runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&ms1)
	o1 := make([]BoundingBoxOrig, N)
	_ = o1
	runtime.ReadMemStats(&ms2)
	t.Logf("Alloc delta for []BoundingBoxOrig(%d): ~%d bytes (%.2f bytes/elem)",
		N, int64(ms2.Alloc-ms1.Alloc), float64(int64(ms2.Alloc-ms1.Alloc))/N)

	runtime.GC()
	runtime.ReadMemStats(&ms1)
	s1 := make([]BoundingBoxSlim, N)
	_ = s1
	runtime.ReadMemStats(&ms2)
	t.Logf("Alloc delta for []BoundingBoxSlim(%d): ~%d bytes (%.2f bytes/elem)",
		N, int64(ms2.Alloc-ms1.Alloc), float64(int64(ms2.Alloc-ms1.Alloc))/N)

	runtime.GC()
	runtime.ReadMemStats(&ms1)
	c1 := make([]ClientOrig, N)
	_ = c1
	runtime.ReadMemStats(&ms2)
	t.Logf("Alloc delta for []ClientOrig(%d):      ~%d bytes (%.2f bytes/elem)",
		N, int64(ms2.Alloc-ms1.Alloc), float64(int64(ms2.Alloc-ms1.Alloc))/N)

	runtime.GC()
	runtime.ReadMemStats(&ms1)
	c2 := make([]ClientOpt, N)
	_ = c2
	runtime.ReadMemStats(&ms2)
	t.Logf("Alloc delta for []ClientOpt(%d):       ~%d bytes (%.2f bytes/elem)",
		N, int64(ms2.Alloc-ms1.Alloc), float64(int64(ms2.Alloc-ms1.Alloc))/N)
}

// ---------- Iterate benchmarks (cache friendliness) ----------
func BenchmarkIterate_BBoxOrig(b *testing.B) {
	const N = 1_000_000
	data := makeBBoxOrig(N)
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(BoundingBoxOrig{})) * N)
	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := 0
		for j := 0; j < N; j++ {
			r := data[j]
			s += r.BlockNum + r.ParNum + r.LineNum + r.WordNum
			if r.Confidence > 0.9999999 {
				s++
			}
			if len(r.Word) == 17 {
				s++
			}
			// touch coords
			s += r.Box.Max.X - r.Box.Min.X + r.Box.Max.Y - r.Box.Min.Y
		}
		sum += s
	}
	blackhole = sum
}

func BenchmarkIterate_BBoxSlim(b *testing.B) {
	const N = 1_000_000
	data := makeBBoxSlim(N)
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(BoundingBoxSlim{})) * N)
	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := 0
		for j := 0; j < N; j++ {
			r := data[j]
			s += int(r.BlockNum + r.ParNum + r.LineNum + r.WordNum)
			if r.Confidence > 0.9999999 {
				s++
			}
			if len(r.Word) == 17 {
				s++
			}
			s += int(r.X1-r.X0) + int(r.Y1-r.Y0)
		}
		sum += s
	}
	blackhole = sum
}

// ---------- Copy benchmarks (struct memmove bandwidth) ----------
func BenchmarkCopy_BBoxOrig(b *testing.B) {
	const N = 200_000
	src := makeBBoxOrig(N)
	dst := make([]BoundingBoxOrig, N)
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(BoundingBoxOrig{})) * N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(dst, src)
	}
	blackhole = dst
}

func BenchmarkCopy_BBoxSlim(b *testing.B) {
	const N = 200_000
	src := makeBBoxSlim(N)
	dst := make([]BoundingBoxSlim, N)
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(BoundingBoxSlim{})) * N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(dst, src)
	}
	blackhole = dst
}

// ---------- Client iteration (mostly shows size/cache effects) ----------
func BenchmarkIterate_ClientOrig(b *testing.B) {
	const N = 500_000
	data := makeClientOrig(N)
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(ClientOrig{})) * N)
	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := 0
		for j := 0; j < N; j++ {
			c := data[j]
			if c.Trim {
				s++
			}
			if c.shouldInit {
				s++
			}
			s += len(c.TessdataPrefix)
			s += len(c.Languages)
			s += len(c.Variables)
			s += len(c.ConfigFilePath)
		}
		sum += s
	}
	blackhole = sum
}

func BenchmarkIterate_ClientOpt(b *testing.B) {
	const N = 500_000
	data := makeClientOpt(N)
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(ClientOpt{})) * N)
	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := 0
		for j := 0; j < N; j++ {
			c := data[j]
			if c.Trim {
				s++
			}
			if c.shouldInit {
				s++
			}
			s += len(c.TessdataPrefix)
			s += len(c.Languages)
			s += len(c.Variables)
			s += len(c.ConfigFilePath)
		}
		sum += s
	}
	blackhole = sum
}

// ---------- Pretty-print sizes after benches ----------
func BenchmarkPrintSizes(b *testing.B) {
	b.StopTimer()
	sizes := []struct {
		name string
		size uintptr
	}{
		{"ClientOrig ", unsafe.Sizeof(ClientOrig{})},
		{"ClientOpt  ", unsafe.Sizeof(ClientOpt{})},
		{"BBoxOrig   ", unsafe.Sizeof(BoundingBoxOrig{})},
		{"BBoxSlim   ", unsafe.Sizeof(BoundingBoxSlim{})},
	}
	for _, s := range sizes {
		fmt.Printf("%s sizeof = %d bytes\n", s.name, s.size)
	}
}
