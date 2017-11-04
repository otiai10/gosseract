#####
# This is an example runtime to use "gosseract" package.
# You can just try how it works by using docker command.
# Hit `docker run -it --rm otiai10/gosseract` to check it out.
#####
FROM debian:latest

LABEL maintainer="otiai10<otiai10@gmail.com>"

RUN apt-get update -qq
RUN apt-get install -y libtesseract-dev libleptonica-dev

RUN apt-get install -y git wget golang

ENV GOPATH=/go

RUN go get github.com/otiai10/gosseract

# For testing
RUN go get github.com/otiai10/mint

# Load languages
ARG LANGS="eng,deu,jpn"
RUN IFS="," &&\
for lang in ${LANGS}; do\
  wget -q https://github.com/tesseract-ocr/tessdata/raw/master/${lang}.traineddata -P /usr/local/share/tessdata;\
done
# If you want to add languages, use --build-arg like this
# > docker build . --build-arg LANGS="eng,fra,spa"
