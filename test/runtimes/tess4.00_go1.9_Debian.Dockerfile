FROM debian:stretch

RUN apt-get update -qq
RUN apt-get install -yq \
  g++ \
  autoconf \
  autoconf-archive \
  automake \
  libtool \
  pkg-config \
  libpng-dev \
  libjpeg62-turbo-dev \
  libtiff5-dev \
  zlib1g-dev
RUN apt-get install -yq \
  libicu-dev \
  libpango1.0-dev \
  libcairo2-dev
RUN apt-get install -yq \
  wget \
  git

ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib

# Compile Leptonica
WORKDIR /
RUN mkdir -p /tmp/leptonica \
  && wget -nv https://github.com/DanBloomberg/leptonica/releases/download/1.76.0/leptonica-1.76.0.tar.gz \
  && tar -xzf leptonica-1.76.0.tar.gz -C /tmp/leptonica \
  && mv /tmp/leptonica/* /leptonica
WORKDIR /leptonica
RUN autoreconf -i \
  && ./autobuild \
  && ./configure \
  && make --quiet \
  && make install

# Compile Tesseract
WORKDIR /
RUN mkdir -p /tmp/tesseract \
  && wget -nv https://github.com/tesseract-ocr/tesseract/archive/4.0.0-beta.3.tar.gz \
  && tar -xzf 4.0.0-beta.3.tar.gz -C /tmp/tesseract \
  && mv /tmp/tesseract/* /tesseract
WORKDIR /tesseract
RUN ./autogen.sh \
  && ./configure \
  && make --quiet \
  && make install

# Recover location
WORKDIR /

# Load languages
RUN wget -nv https://github.com/tesseract-ocr/tessdata/raw/master/eng.traineddata -P /usr/local/share/tessdata
RUN wget -nv https://github.com/tesseract-ocr/tessdata/raw/master/jpn.traineddata -P /usr/local/share/tessdata

# Install Go1.9.1
RUN wget -nv https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz \
  && tar -xzf go1.9.1.linux-amd64.tar.gz
ENV GOROOT=/go

# Prepare GOPATH
RUN mkdir /gopath
ENV GOPATH=/gopath
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

# Dependencies for tests
RUN go get github.com/otiai10/mint
RUN go get golang.org/x/net/html

# Mount source code of gosseract project
ADD . ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
