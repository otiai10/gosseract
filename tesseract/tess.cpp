#if __FreeBSD__ >= 10
#include "/usr/local/include/tesseract/baseapi.h"
#include "/usr/local/include/leptonica/allheaders.h"
#else
#include <tesseract/baseapi.h>
#include <leptonica/allheaders.h>
#endif

extern "C" {
    class TessClient {
      private:
        tesseract::TessBaseAPI *api;
        Pix *image;
      public:
        TessClient()
        {
          api = new tesseract::TessBaseAPI();
        }
        TessClient(char *imgPath)
        {
          image = pixRead(imgPath);
        }
        void setImage(char* imgPath)
        {
          image = pixRead(imgPath);
        }
        char* Exec()
        {
          api->SetImage(image);
          char *outText = api->GetUTF8Text();
          pixDestroy(&image);
          api->End();
          return outText;
        }
    };

    char* simple(char* filepath, char* whitelist ,char* languages) {
      char *out;
      tesseract::TessBaseAPI *api = new tesseract::TessBaseAPI();
      // Initialize tesseract-ocr with English, without specifying tessdata path
      if (api->Init(NULL, languages)) {
        fprintf(stderr, "Could not initialize tesseract.\n");
        exit(1);
      }

      Pix *image = pixRead(filepath);
      api->SetImage(image);

      if (strlen(whitelist) != 0) {
        api->SetVariable("tessedit_char_whitelist", whitelist);
      }

      out = api->GetUTF8Text();
      api->End();
      pixDestroy(&image);

      return out;
    }

}/* extern "C" */
