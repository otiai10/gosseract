FROM otiai10/tesseract:4.00.00dev

RUN apt-get update && apt-get install -y git

RUN wget https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz \
  && tar -xzvf go1.9.1.linux-amd64.tar.gz
RUN mv /go /.go
ENV GOROOT=/.go

RUN mkdir /go
ENV GOPATH=/go

ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

RUN go get github.com/otiai10/mint
RUN go get github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

CMD go test -run Must
