FROM archlinux/base:latest

RUN pacman -Sy --noconfirm \
  gcc \
  git \
  go

RUN pacman -Sy --noconfirm \
  wget \
  make \
  autoconf \
  automake \
  pkg-config \
  libpng \
  libjpeg \
  libtiff \
  libtool

ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
ENV LIBLEPT_HEADERSDIR=/usr/local/include

# Leptonica
RUN mkdir -p /tmp/leptonica && cd /tmp/leptonica \
  && wget -nv https://github.com/DanBloomberg/leptonica/archive/1.77.0.tar.gz \
  && tar -xzf 1.77.0.tar.gz -C /tmp/leptonica \
  && cd /tmp/leptonica/leptonica-1.77.0 && mkdir m4 \
  && autoreconf -i \
  && ./autogen.sh \
  && ./configure \
  && make \
  && make install

# Tesseract
RUN mkdir -p /tmp/tesseract && cd /tmp/tesseract \
  && wget -nv https://github.com/tesseract-ocr/tesseract/archive/4.0.0.tar.gz \
  && tar -xzf 4.0.0.tar.gz -C /tmp/tesseract \
  && cd /tmp/tesseract/tesseract-4.0.0 \
  && ./autogen.sh \
  && ./configure \
  && make \
  && make install

# Languages
ENV TESSDATA_PREFIX=/usr/local/share/tessdata
RUN wget -nv https://github.com/tesseract-ocr/tessdata/blob/master/eng.traineddata?raw=true -O ${TESSDATA_PREFIX}/eng.traineddata

RUN tesseract --version && tesseract --list-langs

ENV GO111MODULE=on
ENV GOPATH=${HOME}/go
ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib
ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

ENV TESS_LSTM_DISABLED=1
CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
