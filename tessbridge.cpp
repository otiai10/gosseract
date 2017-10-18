#include <tesseract/baseapi.h>
#include <leptonica/allheaders.h>
#include "tessbridge.h"

TessBaseAPI Create() {
  tesseract::TessBaseAPI * api = new tesseract::TessBaseAPI();
  return (void*)api;
}

void Free(TessBaseAPI a) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  api->End();
  delete api;
}

void Init(TessBaseAPI a, char* tessdataprefix, char* languages) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  api->Init(tessdataprefix, languages);
}

void SetImage(TessBaseAPI a, char* imagepath) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  Pix *image = pixRead(imagepath);
  api->SetImage(image);
}

char* UTF8Text(TessBaseAPI a) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  return api->GetUTF8Text();
}

const char* Version(TessBaseAPI a) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  const char* v = api->Version();
  return v;
}
