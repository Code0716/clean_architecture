FROM golang:alpine

WORKDIR /web/src

COPY /mysql /mysql

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/Code0716/clean_architecture &&\
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gin-gonic/gin && \
    go get github.com/gin-contrib/cors

