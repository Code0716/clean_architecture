FROM golang:alpine

WORKDIR /web/src

COPY /mysql /mysql

COPY go.mod go.sum ./
RUN go mod download
COPY . .