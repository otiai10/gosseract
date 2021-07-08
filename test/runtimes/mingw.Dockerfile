FROM mwaeckerlin/mingw

RUN apt-get update -y -q

RUN apt-get install -y \
  golang \
  git \
  libtesseract-dev \
  tesseract-ocr-eng

ENV GOPATH=/root/go
ENV GO111MODULE=on

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "./..."]
