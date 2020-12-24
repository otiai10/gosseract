FROM debian:latest

RUN apt-get update -qq
RUN apt-get install -y \
  git \
  golang \
  libtesseract-dev \
  libleptonica-dev \
  tesseract-ocr-eng

ENV GOPATH=/root/go
RUN go get -u github.com/otiai10/mint golang.org/x/net/html

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

RUN tesseract --version

# CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
CMD ["go", "test", "-v", "./..."]
