package gosseract

import (
	"image"
	"runtime"
	"testing"
	"unsafe"
)

// cBoundingBox mirrors the C struct layout we read in production.
// C-ish layout we're simulating:
//
//	struct bounding_box {
//	  char*  word;        // pointer (assume 64-bit here; Go adapts via unsafe.Pointer size)
//	  int    x1, y1, x2, y2;   // model as int32 like common C "int" on 32-bit width
//	  double confidence;       // float64
//	  int    block_num, par_num, line_num, word_num; // int32
//	};
//
// NOTE: This is used only inside this test; it is not exported and not used by the library code.
type cBoundingBox struct {
	word       *byte
	x1, y1     int32
	x2, y2     int32
	confidence float64
	block_num  int32
	par_num    int32
	line_num   int32
	word_num   int32
}

// goCString allocates a NUL-terminated byte buffer containing s and returns a pointer
// to its first byte. We keep the backing slice alive via the returned keep func.
func goCString(s string) (ptr *byte, keep func()) {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return &b[0], func() { runtime.KeepAlive(b) }
}

// cStringToGo reads a NUL-terminated string from p (Go memory) without using cgo.
func cStringToGo(p *byte) string {
	if p == nil {
		return ""
	}
	base := uintptr(unsafe.Pointer(p))
	n := 0
	for {
		b := *(*byte)(unsafe.Pointer(base + uintptr(n)))
		if b == 0 {
			break
		}
		n++
	}
	if n == 0 {
		return ""
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(p)), n))
}

func TestPureGoBoundingBoxesMapping(t *testing.T) {
	// 1) Build a "C-like" array in Go memory.
	arr := make([]cBoundingBox, 3)

	p0, k0 := goCString("hello")
	p1, k1 := goCString("world")
	p2, k2 := goCString("OCR")

	arr[0] = cBoundingBox{
		word: p0,
		x1:   10, y1: 20, x2: 110, y2: 120,
		confidence: 0.987654321,
		block_num:  1, par_num: 2, line_num: 3, word_num: 4,
	}
	arr[1] = cBoundingBox{
		word: p1,
		x1:   -5, y1: 0, x2: 5, y2: 10,
		confidence: 0.5,
		block_num:  0, par_num: 1, line_num: 1, word_num: 2,
	}
	arr[2] = cBoundingBox{
		word: p2,
		x1:   32760, y1: 32761, x2: 40000, y2: 50000,
		confidence: 1.0,
		block_num:  9, par_num: 8, line_num: 7, word_num: 6,
	}

	// 2) Map to Go BoundingBox using the same pointer-stepping style as production code.
	base := unsafe.Pointer(&arr[0])
	step := unsafe.Sizeof(cBoundingBox{})
	length := len(arr)

	out := make([]BoundingBox, 0, length)
	for i := 0; i < length; i++ {
		ptr := unsafe.Pointer(uintptr(base) + uintptr(i)*step)
		box := (*cBoundingBox)(ptr)

		out = append(out, BoundingBox{
			Word:       cStringToGo(box.word),
			X0:         int32(box.x1),
			Y0:         int32(box.y1),
			X1:         int32(box.x2),
			Y1:         int32(box.y2),
			Confidence: float32(box.confidence),
			BlockNum:   int32(box.block_num),
			ParNum:     int32(box.par_num),
			LineNum:    int32(box.line_num),
			WordNum:    int32(box.word_num),
		})
	}

	// 3) Assert values (including Rect() helper).
	want := []struct {
		word       string
		rect       image.Rectangle
		conf       float32
		block, par int32
		line, wn   int32
	}{
		{"hello", image.Rect(10, 20, 110, 120), float32(0.987654321), 1, 2, 3, 4},
		{"world", image.Rect(-5, 0, 5, 10), float32(0.5), 0, 1, 1, 2},
		{"OCR", image.Rect(32760, 32761, 40000, 50000), float32(1.0), 9, 8, 7, 6},
	}
	if len(out) != len(want) {
		t.Fatalf("expected %d boxes, got %d", len(want), len(out))
	}
	const eps = 1e-6
	for i := range want {
		g := out[i]
		w := want[i]
		if g.Word != w.word {
			t.Errorf("[%d] word: got %q want %q", i, g.Word, w.word)
		}
		if g.Rect() != w.rect {
			t.Errorf("[%d] rect: got %+v want %+v", i, g.Rect(), w.rect)
		}
		d := float64(g.Confidence) - float64(w.conf)
		if d > eps || d < -eps {
			t.Errorf("[%d] conf: got %v want %v (|diff|>%g)", i, g.Confidence, w.conf, eps)
		}
		if g.BlockNum != w.block {
			t.Errorf("[%d] block: got %d want %d", i, g.BlockNum, w.block)
		}
		if g.ParNum != w.par {
			t.Errorf("[%d] par: got %d want %d", i, g.ParNum, w.par)
		}
		if g.LineNum != w.line {
			t.Errorf("[%d] line: got %d want %d", i, g.LineNum, w.line)
		}
		if g.WordNum != w.wn {
			t.Errorf("[%d] wordnum: got %d want %d", i, g.WordNum, w.wn)
		}
	}

	// Keep the backing string buffers alive until after use.
	k0()
	k1()
	k2()
}
