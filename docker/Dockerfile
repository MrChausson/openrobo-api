#Dockerfile for OpenRobo API made in golang
FROM golang:1.22.2 AS build

RUN apt update

RUN apt install -y libasound2-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . /app

RUN cd /app/src && go build -v -o /usr/local/bin/openrobo-api

ENTRYPOINT ["/usr/local/bin/openrobo-api"]
