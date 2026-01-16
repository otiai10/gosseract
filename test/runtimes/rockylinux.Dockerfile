FROM rockylinux:9

# Enable EPEL and CRB repositories for tesseract dependencies
RUN dnf install -y epel-release \
    && dnf config-manager --set-enabled crb \
    && dnf install -y \
        tesseract \
        tesseract-devel \
        tesseract-langpack-eng \
        leptonica-devel \
        golang \
        git \
        gcc-c++

ENV GOPATH=/root/go
ENV GO111MODULE=on

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract
RUN go get -t -v ./...

RUN tesseract --version

CMD ["go", "test", "-v", "./..."]
