#ifdef __cplusplus
extern "C" {
#endif

typedef void* TessBaseAPI;
TessBaseAPI Init(void);
void Free(TessBaseAPI);
const char* Version(TessBaseAPI);

#ifdef __cplusplus
}
#endif/* extern "C" */
