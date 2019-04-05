FROM alpine:latest

RUN apk update
RUN apk add \
  g++ \
  git \
  musl-dev \
  go \
  tesseract-ocr-dev

ENV GOPATH=/root/go
RUN go get -u github.com/otiai10/mint golang.org/x/net/html
ADD . ${GOPATH}/src/github.com/otiai10/gosseract

ENV GOSSERACT_CPPSTDERR_NOT_CAPTURED=1
ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
