FROM ubuntu:16.04

ARG TESS="4.00.00dev"
ARG LEPTO="1.74.2"
ARG GO="1.9.1"

RUN apt-get update -qq
RUN apt-get install -yq \
  git \
  wget \
  make \
  autoconf \
  automake \
  libtool \
  autoconf-archive \
  pkg-config \
  libpng-dev \
  libjpeg-dev \
  libtiff-dev \
  zlib1g-dev \
  libicu-dev \
  libpango1.0-dev \
  libcairo2-dev

ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib

# Compile Leptonica
WORKDIR /
RUN mkdir -p /tmp/leptonica \
  && wget -nv https://github.com/DanBloomberg/leptonica/archive/${LEPTO}.tar.gz \
  && tar -xzf ${LEPTO}.tar.gz -C /tmp/leptonica \
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
  && wget -nv https://github.com/tesseract-ocr/tesseract/archive/${TESS}.tar.gz \
  && tar -xzf ${TESS}.tar.gz -C /tmp/tesseract \
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

ENTRYPOINT go test github.com/otiai10/gosseract
