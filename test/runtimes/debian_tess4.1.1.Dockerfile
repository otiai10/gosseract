FROM debian:bullseye-slim

RUN apt-get update -qq
RUN apt-get install -y \
  git \
  golang \
  libtesseract-dev=4.1.1-2.1 \
  tesseract-ocr-eng

ENV GOPATH=/root/go

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

RUN tesseract --version

CMD ["go", "test", "-v", "./..."]
