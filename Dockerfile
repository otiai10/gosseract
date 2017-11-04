#####
# This is a working example of setting up tesseract/gosseract,
# and also works as an example runtime to use gosseract package.
# You can just hit `docker run -it --rm otiai10/gosseract`
# to try and check it out!
#####
FROM golang:latest
LABEL maintainer="otiai10<otiai10@gmail.com>"

RUN apt-get update -qq

# You need librariy files and headers of tesseract and leptonica.
# When you miss these or LD_LIBRARY_PATH is not set to them,
# you would face an error: "tesseract/baseapi.h: No such file or directory"
RUN apt-get install -y libtesseract-dev libleptonica-dev

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr

RUN go get github.com/otiai10/gosseract

# Load languages
ARG LANGS="eng,deu,jpn"
RUN IFS="," && for lang in ${LANGS}; do apt-get install tesseract-ocr-${lang}; done
# If you want to add languages, use --build-arg like this
# > docker build . --build-arg LANGS="eng,fra,spa"

# "mint" is just for minimum testing
RUN go get github.com/otiai10/mint
RUN cd ${GOPATH}/src/github.com/otiai10/gosseract && go test
