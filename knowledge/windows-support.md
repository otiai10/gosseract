# Windows Support for gosseract

This document describes how Windows support was implemented for gosseract and the key learnings from the development process.

## Overview

gosseract is a Go wrapper for Tesseract OCR that uses CGO to interface with the C API. Supporting Windows required solving several challenges related to CGO compilation, library linking, and runtime DLL discovery.

## The Challenge

Before this implementation, Windows support was blocked by issue #262. The main pain points were:

1. **No pre-built Tesseract libraries**: Users had to compile Tesseract from source
2. **MSVC vs MinGW incompatibility**: vcpkg builds Tesseract with MSVC, but Go's CGO on Windows uses MinGW (GCC)
3. **Complex build setup**: Multiple environment variables and paths needed configuration

## Solution Architecture

### 1. Use Tesseract C API Instead of C++ API

The original `tessbridge.cpp` used the C++ API which has name mangling issues between MSVC and MinGW. We rewrote it as `tessbridge.c` using the pure C API (`tesseract/capi.h`), which has stable ABI across compilers.

**Key file**: `tessbridge.c`

### 2. Create MinGW Import Libraries from MSVC DLLs

vcpkg installs MSVC-compiled DLLs (`.dll`) and import libraries (`.lib`). MinGW's linker needs `.a` import libraries. We use `gendef` and `dlltool` (included with MinGW) to create them:

```bash
# Generate .def file from DLL
gendef tesseract55.dll

# Create MinGW import library
dlltool -d tesseract55.def -l libtesseract55.a -D tesseract55.dll
```

### 3. Explicit CGO Flags (No pkg-config)

Windows doesn't have reliable pkg-config support. Instead of:
```go
// #cgo pkg-config: tesseract lept
```

We use explicit flags in `preprocessflags_windows.go`:
```go
// #cgo CFLAGS: -Wno-unused-result
// #cgo LDFLAGS: -ltesseract -lleptonica
```

The actual include/library paths are provided via environment variables:
- `CGO_CFLAGS`: `-IC:/vcpkg/installed/x64-windows/include`
- `CGO_LDFLAGS`: `-LC:/vcpkg/installed/x64-windows/lib`

## Key Learnings

### 1. CGO_ENABLED Must Be Explicitly Set

In GitHub Actions with bash shell on Windows, `CGO_ENABLED` defaults to `0` even though Go detects a C compiler. Always set it explicitly:

```yaml
env:
  CGO_ENABLED: "1"
```

### 2. CC/CXX Must Be Full Paths

When using bash shell on Windows, the `PATH` environment variable in the YAML `env:` section doesn't work as expected for Go's CGO. The C compiler must be specified with full path:

```yaml
env:
  CC: C:/mingw64/bin/gcc.exe
  CXX: C:/mingw64/bin/g++.exe
```

### 3. DLL Discovery at Runtime

Even after successful compilation, tests will fail with `exit status 0xc0000135` (STATUS_DLL_NOT_FOUND) if DLLs aren't in PATH at runtime. Export PATH in the script:

```bash
export PATH="/c/mingw64/bin:/c/vcpkg/installed/x64-windows/bin:$PATH"
```

### 4. Library Symlinks Need Careful Glob Patterns

When creating symlinks like `libtesseract.a -> libtesseract55.a`, be careful with glob patterns. If the symlink already exists, `libtesseract*.a` will match both files:

```bash
# BAD: matches libtesseract.a too
for f in libtesseract*.a; do ln -sf "$f" libtesseract.a; done

# GOOD: only matches versioned files
for f in libtesseract[0-9]*.a; do ln -sf "$f" libtesseract.a; done
```

### 5. Tesseract Needs Language Data

Tests require `eng.traineddata`. Download it and set `TESSDATA_PREFIX`:

```yaml
- run: |
    mkdir -p /c/tessdata
    curl -sL -o /c/tessdata/eng.traineddata \
      https://github.com/tesseract-ocr/tessdata/raw/main/eng.traineddata

- env:
    TESSDATA_PREFIX: C:/tessdata
```

### 6. vcpkg Caching Dramatically Improves Build Time

- **Without cache**: ~22 minutes (compiles Tesseract and all dependencies)
- **With cache**: ~4 minutes

Cache these directories:
```yaml
path: |
  C:/vcpkg/installed
  C:/vcpkg/packages
  C:/vcpkg/buildtrees
```

## File Structure

```
gosseract/
├── tessbridge.c              # C API bridge (not C++)
├── tessbridge.h              # Header file
├── preprocessflags_windows.go # Windows-specific CGO flags
├── preprocessflags_darwin.go  # macOS-specific CGO flags
├── preprocessflags_x.go       # Linux/other platforms
└── .github/workflows/
    └── windows-ci.yml         # Windows CI workflow
```

## CI Workflow Summary

The Windows CI (`windows-ci.yml`) performs these steps:

1. **Checkout** - Get the code
2. **Setup Go** - Install Go 1.24+
3. **Cache vcpkg** - Restore cached packages if available
4. **Install Tesseract** - `vcpkg install tesseract:x64-windows`
5. **Create import libraries** - Generate `.a` files from DLLs
6. **Download tessdata** - Get `eng.traineddata`
7. **Run tests** - With proper environment variables

## For End Users

Users who want to build gosseract on Windows locally need:

1. **MinGW-w64** (for GCC compiler)
2. **vcpkg** with Tesseract installed
3. Environment variables set:
   ```bash
   export CGO_ENABLED=1
   export CC=/c/mingw64/bin/gcc.exe
   export CGO_CFLAGS="-IC:/vcpkg/installed/x64-windows/include"
   export CGO_LDFLAGS="-LC:/vcpkg/installed/x64-windows/lib"
   export TESSDATA_PREFIX="C:/tessdata"
   export PATH="/c/mingw64/bin:/c/vcpkg/installed/x64-windows/bin:$PATH"
   ```

4. MinGW import libraries (`.a` files) created from vcpkg's DLLs

## References

- [Issue #262](https://github.com/otiai10/gosseract/issues/262) - Original Windows support issue
- [Tesseract C API](https://tesseract-ocr.github.io/tessapi/5.x/capi_8h.html)
- [vcpkg](https://vcpkg.io/) - C++ package manager
- [MinGW-w64](https://www.mingw-w64.org/) - GCC for Windows

## Timeline

- **2026-01-16**: Initial implementation completed
- Commits on `windows/ci` branch document the iterative debugging process
