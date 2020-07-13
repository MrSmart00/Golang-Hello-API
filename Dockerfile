FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go get -u github.com/labstack/echo/...
