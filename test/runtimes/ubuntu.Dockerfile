FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -qq \
    && apt-get install -yq \
      git \
      golang \
      libtesseract-dev \
      libleptonica-dev

# Load languages
RUN apt-get install -y \
  tesseract-ocr-eng

ENV GOPATH=/root/go
ENV GO111MODULE=on

# Mount source code of gosseract project
ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "./..."]
