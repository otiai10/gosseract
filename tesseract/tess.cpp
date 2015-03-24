#include <tesseract/baseapi.h>
#include <leptonica/allheaders.h>

extern "C" {
    int hoge() {
        char *outText;

        tesseract::TessBaseAPI *api = new tesseract::TessBaseAPI();
        // Initialize tesseract-ocr with English, without specifying tessdata path
        if (api->Init(NULL, "eng")) {
            fprintf(stderr, "Could not initialize tesseract.\n");
            exit(1);
        }

        // Open input image with leptonica library
        // Pix *image = pixRead("/usr/src/tesseract-3.02/phototest.tif");
        Pix *image = pixRead("sample.png");
        api->SetImage(image);
        // Get OCR result
        outText = api->GetUTF8Text();
        printf("OCR output:\n%s", outText);

        // Destroy used object and release memory
        api->End();
        delete [] outText;
        pixDestroy(&image);

        return 0;
    }

    char* fuga(char* filepath) {
      char *out;
      tesseract::TessBaseAPI *api = new tesseract::TessBaseAPI();
      // Initialize tesseract-ocr with English, without specifying tessdata path
      if (api->Init(NULL, "eng")) {
        fprintf(stderr, "Could not initialize tesseract.\n");
        exit(1);
      }

      Pix *image = pixRead(filepath);
      api->SetImage(image);

      out = api->GetUTF8Text();
      api->End();
      pixDestroy(&image);

      return out;
    }

}/* extern "C" */
