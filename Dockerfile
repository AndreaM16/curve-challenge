FROM golang:1.9

RUN mkdir -p /go/src/github.com/andream16/curve-challenge
WORKDIR /go/src/github.com/andream16/curve-challenge

ADD . /go/src/github.com/andream16/curve-challenge

RUN go get -v