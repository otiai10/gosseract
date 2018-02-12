FROM mwaeckerlin/mingw

RUN apt-get update
RUN apt-get install -y git

# Packages
RUN apt-get install -y \
  libtesseract-dev \
  libleptonica-dev \
  tesseract-ocr-eng

# Golang itself
RUN wget -nv https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz \
 && tar -xzf go1.9.2.linux-amd64.tar.gz -C /
ENV GOROOT=/go
RUN ls -la ${GOROOT}/bin
RUN mkdir /gopath
ENV GOPATH=/gopath
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

# Dependencies for tests
RUN go get github.com/otiai10/mint
RUN go get golang.org/x/net/html

# Mount source code of gosseract project
# instead of `go get github.com/otiai10/gosseract`
ADD . ${GOPATH}/src/github.com/otiai10/gosseract

ENTRYPOINT go test github.com/otiai10/gosseract
