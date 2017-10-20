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

void SetVariable(TessBaseAPI a, char* name, char* value) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  api->SetVariable(name, value);
}

void SetImage(TessBaseAPI a, char* imagepath) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  Pix *image = pixRead(imagepath);
  api->SetImage(image);
}

void SetPageSegMode(TessBaseAPI a, int m) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  tesseract::PageSegMode mode = (tesseract::PageSegMode)m;
  api->SetPageSegMode(mode);
}

int GetPageSegMode(TessBaseAPI a) {
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  return api->GetPageSegMode();
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
