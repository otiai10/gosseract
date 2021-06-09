FROM fedora

RUN yum update -y -q

RUN yum install -y -q \
  go \
  gcc-c++ \
  tesseract-devel

ENV GOPATH=/root/go

RUN go get -u github.com/otiai10/mint golang.org/x/net/html
ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

CMD ["go", "test", "-v", "./..."]
