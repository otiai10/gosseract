FROM archlinux/base:latest

RUN pacman -Sy -q --noconfirm \
  gcc \
  glibc \
  git \
  tesseract \
  tesseract-data-eng \
  go

ENV TESSDATA_PREFIX=/usr/share/tessdata
ENV GOPATH=${HOME}/go
ENV GO111MODULE=on

# Dependencies for tests
RUN go get -u github.com/otiai10/mint golang.org/x/net/html

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

RUN tesseract --version
CMD ["go", "test", "-v", "./..."]
