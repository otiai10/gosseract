FROM mwaeckerlin/mingw

RUN apt-get update -y -q

RUN apt-get install -y \
  wget \
  git \
  gcc \
  g++ \
  libtesseract-dev \
  tesseract-ocr-eng

# Install Go 1.21 manually (apt golang is too old)
RUN wget -q https://go.dev/dl/go1.21.13.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf go1.21.13.linux-amd64.tar.gz \
  && rm go1.21.13.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH=/root/go
ENV GO111MODULE=on

ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "./..."]
