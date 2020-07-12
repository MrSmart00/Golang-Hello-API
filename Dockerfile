FROM golang:latest

RUN mkdir /go/src/api

WORKDIR /go/src/api

ADD . /go/src/api

RUN go get github.com/gin-gonic/gin
