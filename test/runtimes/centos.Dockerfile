FROM centos:latest

RUN yum update -y
RUN yum install -y yum-plugin-ovl epel-release
RUN yum install -y \
  golang \
  make \
  gcc-c++ \
  wget \
  git \
  tar \
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

# Leptonica
RUN mkdir -p /tmp/leptonica \
  && wget -nv https://github.com/DanBloomberg/leptonica/archive/1.74.4.tar.gz \
  && tar -xzf 1.74.4.tar.gz -C /tmp/leptonica \
  && mv /tmp/leptonica/* /leptonica && cd /leptonica && mkdir m4 \
  && autoreconf -i \
  && chmod a+x ./autobuild && ./autobuild \
  && chmod a+x ./configure && ./configure \
  && make && make install

ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
ENV LIBLEPT_HEADERSDIR=/usr/local/include

# Tesseract
RUN mkdir -p /tmp/tesseract \
  && wget -nv https://github.com/tesseract-ocr/tesseract/archive/3.05.02.tar.gz \
  && tar -xzf 3.05.02.tar.gz -C /tmp/tesseract \
  && mv /tmp/tesseract/* /tesseract && cd /tesseract \
  && ./autogen.sh && ./configure \
  && make && make install

ENV TESSDATA_PREFIX=/usr/share/tesseract
RUN mkdir -p ${TESSDATA_PREFIX}/tessdata
RUN wget -nv https://github.com/tesseract-ocr/tessdata/raw/3.04.00/eng.traineddata \
  -O ${TESSDATA_PREFIX}/tessdata/eng.traineddata

ENV GOPATH=/root/go
# Dependencies for tests
RUN go get github.com/otiai10/mint golang.org/x/net/html
# Mount source code of gosseract project
ADD . ${GOPATH}/src/github.com/otiai10/gosseract

CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
