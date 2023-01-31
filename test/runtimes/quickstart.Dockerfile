FROM golang:latest

# Update registry and install tesseract and dependencies
RUN apt-get update -qq \
    && apt-get install -y \
      libtesseract-dev \
      libleptonica-dev \
      tesseract-ocr-eng

WORKDIR $GOPATH/src/your-project
RUN go mod init

RUN go get -u -v -t github.com/otiai10/gosseract/v2

# Test it!
CMD ["go", "test", "-v", "github.com/otiai10/gosseract/v2"]
