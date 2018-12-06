FROM base/archlinux:latest

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
  libtool

ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
ENV LIBLEPT_HEADERSDIR=/usr/local/include

# Leptonica
RUN mkdir -p /tmp/leptonica && cd /tmp/leptonica \
  && wget https://github.com/DanBloomberg/leptonica/archive/1.76.0.tar.gz \
  && tar -xzf 1.76.0.tar.gz -C /tmp/leptonica \
  && cd /tmp/leptonica/leptonica-1.76.0 && mkdir m4 \
  && autoreconf -i \
  && ./autobuild \
  && ./configure \
  && make \
  && make install

# Tesseract
RUN mkdir -p /tmp/tesseract && cd /tmp/tesseract \
  && wget https://github.com/tesseract-ocr/tesseract/archive/3.05.02.tar.gz \
  && tar -xzf 3.05.02.tar.gz -C /tmp/tesseract \
  && cd /tmp/tesseract/tesseract-3.05.02 \
  && ./autogen.sh \
  && ./configure \
  && make \
  && make install

# Languages
RUN wget https://github.com/tesseract-ocr/tessdata/blob/master/eng.traineddata?raw=true -O /usr/local/share/tessdata/eng.traineddata

ENV TESSDATA_PREFIX=/usr/local/share/tessdata
# RUN tesseract --version && tesseract --list-langs

ENV GOPATH=${HOME}/go
ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib
ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

# Dependencies for tests
RUN go get -u -v github.com/otiai10/mint golang.org/x/net/html

CMD ["go", "test", "-v", "github.com/otiai10/gosseract"]
