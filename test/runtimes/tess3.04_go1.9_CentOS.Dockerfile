FROM centos:latest

RUN yum update -y -q
RUN yum install -y -q \
  gcc-c++ \
  autoconf \
  autoconf-archive \
  automake \
  make \
  libtool \
  libjpeg-devel \
  libpng-devel \
  libtiff-devel \
  zlib-devel
RUN yum install -y -q \
  libicu-devel \
  libpango1.0-dev \
  libcairo-dev
RUN yum install -y -q \
  wget \
  git

ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib
ENV TESSDATA_PREFIX=/usr/local/share

# Compile Leptonica
# WORKDIR /
# RUN mkdir -p /tmp/leptonica \
#   && wget -nv https://github.com/DanBloomberg/leptonica/releases/download/1.76.0/leptonica-1.76.0.tar.gz \
#   && tar -xzf leptonica-1.76.0.tar.gz -C /tmp/leptonica \
#   && mv /tmp/leptonica/* /leptonica
# WORKDIR /leptonica
# RUN autoreconf -i 2>/dev/null \
#   && ./autobuild --silent \
#   && ./configure --enable-silent-rules \
#   && make 1>/dev/null \
#   && make install --silent
RUN yum install -y epel-release
RUN yum install -y leptonica-devel

# Compile Tesseract
# WORKDIR /
# RUN mkdir -p /tmp/tesseract \
#   && wget -nv https://github.com/tesseract-ocr/tesseract/archive/4.0.0-beta.3.tar.gz \
#   && tar -xzf 4.0.0-beta.3.tar.gz -C /tmp/tesseract \
#   && mv /tmp/tesseract/* /tesseract
# WORKDIR /tesseract
# RUN ./autogen.sh \
#   && ./configure --enable-silent-rules \
#   && make 1>/dev/null \
#   && make install --silent
RUN yum install -y tesseract-devel

# Load languages
RUN wget -nv https://github.com/tesseract-ocr/tessdata/raw/master/eng.traineddata -P /usr/local/share/tessdata

# Recover location
WORKDIR /

# Install Go
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
CMD ["go", "test", "github.com/otiai10/gosseract"]
