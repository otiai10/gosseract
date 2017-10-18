#include <tesseract/baseapi.h>
#include <leptonica/allheaders.h>
#include "tessbridge.h"

TessBaseAPI Init()
{
  tesseract::TessBaseAPI * api = new tesseract::TessBaseAPI();
  char* lang;
  api->Init(NULL, lang);
  return (void*)api;
}

void Free(TessBaseAPI a)
{
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  api->End();
  delete api;
}

const char* Version(TessBaseAPI a)
{
  tesseract::TessBaseAPI * api = (tesseract::TessBaseAPI*)a;
  const char* v = api->Version();
  return v;
}
