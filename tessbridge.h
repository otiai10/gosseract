#ifdef __cplusplus
extern "C" {
#endif

typedef void* TessBaseAPI;
TessBaseAPI Create(void);
void Free(TessBaseAPI);
void Init(TessBaseAPI, char*, char*);
void InitConfig(TessBaseAPI, char*, char*, char*);
void SetVariable(TessBaseAPI, char*, char*);
void SetImage(TessBaseAPI, char*);
void SetPageSegMode(TessBaseAPI, int);
int GetPageSegMode(TessBaseAPI);
char* UTF8Text(TessBaseAPI);
char* HOCRText(TessBaseAPI);
const char* Version(TessBaseAPI);

#ifdef __cplusplus
}
#endif/* extern "C" */
