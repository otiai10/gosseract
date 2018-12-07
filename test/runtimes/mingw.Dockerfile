FROM mwaeckerlin/mingw

RUN apt-get update -y -q

# Packages
RUN apt-get install -y \
  golang \
  git \
  libtesseract-dev \
  tesseract-ocr-eng

ENV GOPATH=/root/go

# Dependencies for tests
RUN go get github.com/otiai10/mint golang.org/x/net/html

ADD . ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
