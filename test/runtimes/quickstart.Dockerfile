FROM golang:latest

# Update registry
RUN apt-get update -qq

# Install tesseract and dependencies
RUN apt-get install -y \
  libtesseract-dev \
  libleptonica-dev \
  tesseract-ocr-eng

# Get go packages
RUN go get -t github.com/otiai10/gosseract

# Test it!
CMD go test github.com/otiai10/gosseract
