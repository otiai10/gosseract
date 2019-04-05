FROM archlinux/base:latest

RUN pacman -Sy -q --noconfirm \
  gcc \
  git \
  tesseract \
  tesseract-data-eng \
  go

ENV TESSDATA_PREFIX=/usr/share/tessdata
ENV GOPATH=${HOME}/go

# Dependencies for tests
RUN go get -u github.com/otiai10/mint golang.org/x/net/html

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
