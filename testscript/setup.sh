
# Prepare dependencies
apt-get update
apt-get install -y \
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

export LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib

# Compile Leptonica
cd /
mkdir -p /tmp/leptonica \
  && wget https://github.com/DanBloomberg/leptonica/archive/${LEPTONICA}.tar.gz \
  && tar -xzvf ${LEPTONICA}.tar.gz -C /tmp/leptonica \
  && mv /tmp/leptonica/* /leptonica
cd /leptonica
autoreconf -i \
  && ./autobuild \
  && ./configure \
  && make \
  && make install

# Compile Tesseract
cd /
mkdir -p /tmp/tesseract \
  && wget https://github.com/tesseract-ocr/tesseract/archive/${TESSERACT}.tar.gz \
  && tar -xzvf ${TESSERACT}.tar.gz -C /tmp/tesseract \
  && mv /tmp/tesseract/* /tesseract
cd /tesseract
./autogen.sh \
  && ./configure \
  && make \
  && make install

# Recover location
cd /

# Load languages
wget https://github.com/tesseract-ocr/tessdata/raw/master/eng.traineddata -P /usr/local/share/tessdata
wget https://github.com/tesseract-ocr/tessdata/raw/master/deu.traineddata -P /usr/local/share/tessdata
wget https://github.com/tesseract-ocr/tessdata/raw/master/fra.traineddata -P /usr/local/share/tessdata
wget https://github.com/tesseract-ocr/tessdata/raw/master/spa.traineddata -P /usr/local/share/tessdata
wget https://github.com/tesseract-ocr/tessdata/raw/master/jpn.traineddata -P /usr/local/share/tessdata
