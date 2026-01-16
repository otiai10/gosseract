# Windows Troubleshooting Guide

Common errors encountered when building gosseract on Windows and their solutions.

## Error: "undefined: Version, NewClient, etc."

**Symptom:**
```
.\all_test.go:23:23: undefined: Version
.\all_test.go:38:12: undefined: NewClient
```

**Cause:** CGO is disabled or failing silently. Files with `import "C"` are being skipped.

**Solutions:**
1. Ensure `CGO_ENABLED=1` is set
2. Verify GCC is installed and in PATH
3. Check that `CC` points to a valid compiler

**Debug:**
```bash
go env CGO_ENABLED  # Should be 1
go env CC           # Should show gcc path
```

---

## Error: "cgo: C compiler 'gcc' not found"

**Symptom:**
```
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
```

**Cause:** Go can't find GCC. Even if `which gcc` works in bash, Go uses Windows APIs to find the compiler.

**Solution:** Set `CC` to the full path:
```bash
export CC=C:/mingw64/bin/gcc.exe
```

---

## Error: "exit status 0xc0000135"

**Symptom:**
```
exit status 0xc0000135
FAIL    github.com/otiai10/gosseract/v2    0.028s
```

**Cause:** This is `STATUS_DLL_NOT_FOUND`. The test binary compiled successfully but can't find required DLLs at runtime.

**Solution:** Ensure DLL directories are in PATH when running tests:
```bash
export PATH="/c/vcpkg/installed/x64-windows/bin:$PATH"
```

---

## Error: "Error opening data file ./eng.traineddata"

**Symptom:**
```
Error opening data file ./eng.traineddata
Please make sure the TESSDATA_PREFIX environment variable is set to your "tessdata" directory.
Failed loading language 'eng'
```

**Cause:** Tesseract can't find language data files.

**Solution:**
1. Download language data:
   ```bash
   curl -L -o /c/tessdata/eng.traineddata \
     https://github.com/tesseract-ocr/tessdata/raw/main/eng.traineddata
   ```
2. Set `TESSDATA_PREFIX`:
   ```bash
   export TESSDATA_PREFIX=C:/tessdata
   ```

---

## Error: "undefined reference to `TessBaseAPICreate`"

**Symptom:**
```
undefined reference to `TessBaseAPICreate'
undefined reference to `TessBaseAPIInit3'
```

**Cause:** Linker can't find Tesseract library.

**Solutions:**
1. Ensure import libraries exist in lib directory:
   ```bash
   ls /c/vcpkg/installed/x64-windows/lib/libtesseract.a
   ```
2. Verify `CGO_LDFLAGS` includes the library path:
   ```bash
   export CGO_LDFLAGS="-LC:/vcpkg/installed/x64-windows/lib"
   ```
3. Check that `preprocessflags_windows.go` has correct LDFLAGS

---

## Error: "ln: 'libtesseract.a' and 'libtesseract.a' are the same file"

**Symptom:**
```
ln: 'libtesseract.a' and 'libtesseract.a' are the same file
```

**Cause:** Glob pattern `libtesseract*.a` matches the symlink itself.

**Solution:** Use more specific glob:
```bash
# Instead of libtesseract*.a
for f in libtesseract[0-9]*.a; do
  ln -sf "$f" libtesseract.a
done
```

---

## Error: "cannot find -ltesseract"

**Symptom:**
```
/usr/bin/ld: cannot find -ltesseract
```

**Cause:** MinGW import library doesn't exist or has wrong name.

**Solution:**
1. Create import library from DLL:
   ```bash
   cd /c/vcpkg/installed/x64-windows/bin
   gendef tesseract55.dll
   dlltool -d tesseract55.def -l libtesseract55.a -D tesseract55.dll
   mv libtesseract55.a ../lib/
   ```
2. Create symlink with standard name:
   ```bash
   cd /c/vcpkg/installed/x64-windows/lib
   ln -sf libtesseract55.a libtesseract.a
   ```

---

## Debugging Tips

### Check CGO Environment
```bash
go env | grep -E '(CGO|CC|CXX)'
```

### Verbose Build Output
```bash
go build -x -v . 2>&1 | head -50
```

### List Available Libraries
```bash
ls -la /c/vcpkg/installed/x64-windows/lib/*.a
ls -la /c/vcpkg/installed/x64-windows/lib/*.lib
```

### Check Header Files
```bash
ls /c/vcpkg/installed/x64-windows/include/tesseract/
ls /c/vcpkg/installed/x64-windows/include/leptonica/
```

### Verify DLLs Exist
```bash
ls /c/vcpkg/installed/x64-windows/bin/*.dll | grep -E '(tesseract|leptonica)'
```

### Test GCC Works
```bash
$CC --version
```
