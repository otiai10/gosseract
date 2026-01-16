#if __FreeBSD__ >= 10
#include "/usr/local/include/leptonica/allheaders.h"
#include "/usr/local/include/tesseract/capi.h"
#else
#include <leptonica/allheaders.h>
#include <tesseract/capi.h>
#endif

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifdef _WIN32
#include <io.h>
#define dup _dup
#define dup2 _dup2
#define close _close
#define STDERR_FILENO 2
#define NULL_DEVICE "NUL"
#else
#include <unistd.h>
#define NULL_DEVICE "/dev/null"
#endif

/*
 * Local type definitions matching tessbridge.h
 * We don't include tessbridge.h to avoid TessBaseAPI typedef conflict
 */
typedef void* TessHandle;
typedef void* PixImage;

struct bounding_box {
    int x1, y1, x2, y2;
    char* word;
    float confidence;
    int block_num, par_num, line_num, word_num;
};

struct bounding_boxes {
    int length;
    struct bounding_box* boxes;
};

TessHandle Create(void) {
    return (TessHandle)TessBaseAPICreate();
}

void Free(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    if (api != NULL) {
        TessBaseAPIEnd(api);
        TessBaseAPIDelete(api);
    }
}

void Clear(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    if (api != NULL) {
        TessBaseAPIClear(api);
    }
}

void ClearPersistentCache(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    TessBaseAPIClearPersistentCache(api);
}

int Init(TessHandle a, char* tessdataprefix, char* languages, char* configfilepath, char* errbuf) {
    TessBaseAPI* api = (TessBaseAPI*)a;

    /* Redirect STDERR to given buffer */
    fflush(stderr);
    int original_stderr;
    original_stderr = dup(STDERR_FILENO);
    (void)freopen(NULL_DEVICE, "a", stderr);
    setbuf(stderr, errbuf);

    int ret;
    if (configfilepath != NULL) {
        char* configs[] = {configfilepath};
        int configs_size = 1;
        ret = TessBaseAPIInit1(api, tessdataprefix, languages, OEM_DEFAULT, configs, configs_size);
    } else {
        ret = TessBaseAPIInit3(api, tessdataprefix, languages);
    }

    /* Restore default stderr */
    (void)freopen(NULL_DEVICE, "a", stderr);
    dup2(original_stderr, STDERR_FILENO);
    close(original_stderr);
    setbuf(stderr, NULL);

    return ret;
}

bool SetVariable(TessHandle a, char* name, char* value) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    return TessBaseAPISetVariable(api, name, value) ? true : false;
}

void SetPixImage(TessHandle a, PixImage pix) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    struct Pix* image = (struct Pix*)pix;
    TessBaseAPISetImage2(api, image);
    if (TessBaseAPIGetSourceYResolution(api) < 70) {
        TessBaseAPISetSourceResolution(api, 70);
    }
}

void SetPageSegMode(TessHandle a, int m) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    TessBaseAPISetPageSegMode(api, (TessPageSegMode)m);
}

int GetPageSegMode(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    return (int)TessBaseAPIGetPageSegMode(api);
}

char* UTF8Text(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    return TessBaseAPIGetUTF8Text(api);
}

char* HOCRText(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    return TessBaseAPIGetHOCRText(api, 0);
}

struct bounding_boxes* GetBoundingBoxesVerbose(TessHandle a) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    struct bounding_boxes* box_array;
    box_array = (struct bounding_boxes*)malloc(sizeof(struct bounding_boxes));
    /* linearly resize boxes array */
    int realloc_threshold = 900;
    int realloc_raise = 1000;
    int capacity = 1000;
    box_array->boxes = (struct bounding_box*)malloc(capacity * sizeof(struct bounding_box));
    box_array->length = 0;
    TessBaseAPIRecognize(api, NULL);
    int block_num = 0;
    int par_num = 0;
    int line_num = 0;
    int word_num = 0;

    TessResultIterator* res_it = TessBaseAPIGetIterator(api);
    if (res_it == NULL) {
        return box_array;
    }

    TessPageIterator* page_it = TessResultIteratorGetPageIterator(res_it);

    do {
        /* Check if there's text at the word level */
        char* text = TessResultIteratorGetUTF8Text(res_it, RIL_WORD);
        if (text == NULL) {
            /* Skip empty words */
            continue;
        }

        /* Add rows for any new block/paragraph/textline. */
        if (TessPageIteratorIsAtBeginningOf(page_it, RIL_BLOCK)) {
            block_num++;
            par_num = 0;
            line_num = 0;
            word_num = 0;
        }
        if (TessPageIteratorIsAtBeginningOf(page_it, RIL_PARA)) {
            par_num++;
            line_num = 0;
            word_num = 0;
        }
        if (TessPageIteratorIsAtBeginningOf(page_it, RIL_TEXTLINE)) {
            line_num++;
            word_num = 0;
        }
        word_num++;

        if (box_array->length >= realloc_threshold) {
            capacity += realloc_raise;
            box_array->boxes = (struct bounding_box*)realloc(box_array->boxes, capacity * sizeof(struct bounding_box));
            realloc_threshold += realloc_raise;
        }

        box_array->boxes[box_array->length].word = text;
        box_array->boxes[box_array->length].confidence = TessResultIteratorConfidence(res_it, RIL_WORD);
        TessPageIteratorBoundingBox(page_it, RIL_WORD,
                                    &box_array->boxes[box_array->length].x1,
                                    &box_array->boxes[box_array->length].y1,
                                    &box_array->boxes[box_array->length].x2,
                                    &box_array->boxes[box_array->length].y2);

        /* block, para, line, word numbers */
        box_array->boxes[box_array->length].block_num = block_num;
        box_array->boxes[box_array->length].par_num = par_num;
        box_array->boxes[box_array->length].line_num = line_num;
        box_array->boxes[box_array->length].word_num = word_num;

        box_array->length++;
    } while (TessResultIteratorNext(res_it, RIL_WORD));

    TessResultIteratorDelete(res_it);
    return box_array;
}

struct bounding_boxes* GetBoundingBoxes(TessHandle a, int pageIteratorLevel) {
    TessBaseAPI* api = (TessBaseAPI*)a;
    struct bounding_boxes* box_array;
    box_array = (struct bounding_boxes*)malloc(sizeof(struct bounding_boxes));
    /* linearly resize boxes array */
    int realloc_threshold = 900;
    int realloc_raise = 1000;
    int capacity = 1000;
    box_array->boxes = (struct bounding_box*)malloc(capacity * sizeof(struct bounding_box));
    box_array->length = 0;
    TessBaseAPIRecognize(api, NULL);
    TessResultIterator* ri = TessBaseAPIGetIterator(api);
    TessPageIteratorLevel level = (TessPageIteratorLevel)pageIteratorLevel;

    if (ri != NULL) {
        TessPageIterator* page_it = TessResultIteratorGetPageIterator(ri);
        do {
            if (box_array->length >= realloc_threshold) {
                capacity += realloc_raise;
                box_array->boxes = (struct bounding_box*)realloc(box_array->boxes, capacity * sizeof(struct bounding_box));
                realloc_threshold += realloc_raise;
            }
            box_array->boxes[box_array->length].word = TessResultIteratorGetUTF8Text(ri, level);
            box_array->boxes[box_array->length].confidence = TessResultIteratorConfidence(ri, level);
            TessPageIteratorBoundingBox(page_it, level,
                                        &box_array->boxes[box_array->length].x1,
                                        &box_array->boxes[box_array->length].y1,
                                        &box_array->boxes[box_array->length].x2,
                                        &box_array->boxes[box_array->length].y2);
            box_array->length++;
        } while (TessResultIteratorNext(ri, level));
        TessResultIteratorDelete(ri);
    }

    return box_array;
}

const char* Version(TessHandle a) {
    (void)a;  /* unused parameter */
    return TessVersion();
}

PixImage CreatePixImageByFilePath(char* imagepath) {
    struct Pix* image = pixRead(imagepath);
    return (void*)image;
}

PixImage CreatePixImageFromBytes(unsigned char* data, int size) {
    struct Pix* image = pixReadMem(data, (size_t)size);
    return (void*)image;
}

void DestroyPixImage(PixImage pix) {
    struct Pix* img = (struct Pix*)pix;
    pixDestroy(&img);
}

const char* GetDataPath(void) {
    static TessBaseAPI* api = NULL;
    if (api == NULL) {
        api = TessBaseAPICreate();
        TessBaseAPIInit3(api, NULL, NULL);
    }
    return TessBaseAPIGetDatapath(api);
}
