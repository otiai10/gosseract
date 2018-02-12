FROM centos:6.7

ARG TESS="3.02.02"
ARG LEPTO="v1.72"
ARG GO="1.7.6"

RUN yum update -y
RUN yum install -y yum-plugin-ovl
RUN yum install -y \
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
ENV TESSDATA_PREFIX=/usr/local/share/tesseract-ocr

# Compile Leptonica
WORKDIR /
RUN mkdir -p /tmp/leptonica \
  && wget -nv https://github.com/DanBloomberg/leptonica/archive/${LEPTO}.tar.gz \
  && tar -xzf ${LEPTO}.tar.gz -C /tmp/leptonica \
  && mv /tmp/leptonica/* /leptonica
WORKDIR /leptonica

RUN mkdir m4
RUN autoreconf -i
RUN chmod a+x ./autobuild
RUN ./autobuild
RUN chmod a+x ./configure
RUN ./configure --enable-silent-rules
RUN make --silent
RUN make install --silent

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
# {{{ Tesseract3.02 doesn't work with the latest ttraineddata.
# Use https://github.com/tesseract-ocr/tesseract/wiki/Data-Files#data-files-for-version-302 for older tesseract version.
RUN wget -nv https://downloads.sourceforge.net/project/tesseract-ocr-alt/tesseract-ocr-3.02.eng.tar.gz \
  && tar -xzvf ./tesseract-ocr-3.02.eng.tar.gz  \
  && mv tesseract-ocr ${TESSDATA_PREFIX}/
# Instead of
#
#   RUN wget -nv https://github.com/tesseract-ocr/tessdata/raw/master/eng.traineddata -P /usr/local/share/tessdata
#
# Related issue link: https://github.com/otiai10/gosseract/issues/97
# }}}

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
