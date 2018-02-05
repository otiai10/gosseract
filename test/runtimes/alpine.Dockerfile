FROM alpine:latest

RUN apk update
RUN apk add \
  g++ \
  git \
  musl-dev \
  go \
  tesseract-ocr-dev

RUN go get -u -t -v github.com/otiai10/gosseract

ENTRYPOINT go test github.com/otiai10/gosseract
