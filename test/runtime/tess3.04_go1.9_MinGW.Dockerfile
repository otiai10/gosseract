FROM mwaeckerlin/mingw

RUN apt-get update
RUN apt-get install -y git

# Packages
RUN apt-get install -y \
  libtesseract-dev \
  libleptonica-dev \
  tesseract-ocr-eng

# Golang itself
RUN wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz \
 && tar -xzvf go1.9.2.linux-amd64.tar.gz && mv ./go /.go
ENV GOROOT=/.go
RUN mkdir /go
ENV GOPATH=/go
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

# Dependencies for tests
RUN go get github.com/otiai10/mint

# Mount source code of gosseract project
# instead of `go get github.com/otiai10/gosseract`
ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

ENTRYPOINT go test
