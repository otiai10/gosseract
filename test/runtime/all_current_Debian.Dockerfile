FROM debian:latest

# Update registry
RUN apt-get update -qq

# Install tesseract and dependencies
RUN apt-get install -y \
  libtesseract-dev \
  libleptonica-dev \
  tesseract-ocr-eng

# Install Go
RUN apt-get install -y git golang
ENV GOPATH=/go

# Get go packages
RUN go get github.com/otiai10/mint
ADD . ${GOPATH}/src/github.com/otiai10/gosseract
WORKDIR ${GOPATH}/src/github.com/otiai10/gosseract

# Test it!
CMD go test
