package tesseract

/*
#include <stdio.h>

void hoge(char *s) {
  printf("%s", s);
}
*/
import "C"

// AAA ...
func AAA() {
	s := C.CString("ああああああああ")
	C.hoge(s)
}
