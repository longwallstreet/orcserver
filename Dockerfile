FROM golang:1.14

LABEL maintainer="longwallstreet <jsky1989@gmail.com>"

RUN apt-get update && apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr

ADD . $GOPATH/src/github.com/longwallstreet/ocrserver
WORKDIR $GOPATH/src/github.com/longwallstreet/ocrserver

RUN go get ./...

CMD $GOPATH/bin/ocrserver