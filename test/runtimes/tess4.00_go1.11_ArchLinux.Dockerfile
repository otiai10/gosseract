FROM base/archlinux:latest

RUN pacman -Sy --noconfirm \
  gcc \
  git \
  tesseract \
  tesseract-data-eng \
  go

ENV TESSDATA_PREFIX=/usr/share/tessdata
ENV GOPATH=${HOME}/go

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract
RUN go get -t ./...

CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
