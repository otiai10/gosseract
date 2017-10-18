#ifdef __cplusplus
extern "C" {
#endif

typedef void* TessBaseAPI;
TessBaseAPI Create(void);
void Free(TessBaseAPI);
void Init(TessBaseAPI, char*, char*);
void SetImage(TessBaseAPI, char*);
char* UTF8Text(TessBaseAPI);
const char* Version(TessBaseAPI);

#ifdef __cplusplus
}
#endif/* extern "C" */
