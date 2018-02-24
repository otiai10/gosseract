FROM ubuntu:16.04

RUN apt-get update -qq
RUN apt-get install -yq git golang

# Specify repository for tesseract-ocr & libleptonica-dev
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:alex-p/tesseract-ocr && apt-get update -qq

RUN apt-get install -y \
  tesseract-ocr-dev \
  libleptonica-dev

# Load languages
RUN apt-get install -y \
  tesseract-ocr-eng

RUN mkdir /gopath
ENV GOPATH=/gopath
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

# Dependencies for tests
RUN go get github.com/otiai10/mint
RUN go get golang.org/x/net/html

# Mount source code of gosseract project
ADD . ${GOPATH}/src/github.com/otiai10/gosseract

ENTRYPOINT TESS_LSTM=1 go test -run Test github.com/otiai10/gosseract
