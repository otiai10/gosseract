FROM centos:latest

ARG TESS="4.00.00dev"
ARG LEPTO="1.74.2"
ARG GO="1.9.1"

RUN yum update -y -q
RUN yum install -y -q \
  gcc-c++ \
  git \
  wget \
  make \
  autoconf \
  automake \
  libtool \
  libjpeg-devel \
  libpng-devel \
  libtiff-devel \
  libicu-devel \
  libpango1.0-dev \
  libcairo-dev \
  zlib-devel


ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib
ENV TESSDATA_PREFIX=/usr/local/share

# Compile Leptonica
WORKDIR /
RUN mkdir -p /tmp/leptonica \
  && wget -nv https://github.com/DanBloomberg/leptonica/archive/${LEPTO}.tar.gz \
  && tar -xzf ${LEPTO}.tar.gz -C /tmp/leptonica \
  && mv /tmp/leptonica/* /leptonica
WORKDIR /leptonica

RUN autoreconf -i 2>/dev/null \
  && ./autobuild --silent \
  && ./configure --enable-silent-rules \
  && make 1>/dev/null \
  && make install --silent

# Compile Tesseract
WORKDIR /
RUN mkdir -p /tmp/tesseract \
  && wget -nv https://github.com/tesseract-ocr/tesseract/archive/${TESS}.tar.gz \
  && tar -xzf ${TESS}.tar.gz -C /tmp/tesseract \
  && mv /tmp/tesseract/* /tesseract
WORKDIR /tesseract

RUN ./autogen.sh \
  && ./configure --enable-silent-rules \
  && make 1>/dev/null \
  && make install --silent

# Load languages
RUN wget -nv https://github.com/tesseract-ocr/tessdata/raw/master/eng.traineddata -P /usr/local/share/tessdata

# Recover location
WORKDIR /

# Install Go
RUN wget -nv https://storage.googleapis.com/golang/go${GO}.linux-amd64.tar.gz \
  && tar -xzf go${GO}.linux-amd64.tar.gz
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
