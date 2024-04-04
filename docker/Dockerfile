#Dockerfile for Echo API made in golang
FROM golang:1.21-rc-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . /app

RUN cd /app/ && go build ./src -v -o /usr/local/bin/openrobo-api


#FROM gcr.io/distroless/static:debug
FROM golang:alpine

WORKDIR /app

COPY --from=build /usr/local/bin/openrobo-api /app/openrobo-api

ENTRYPOINT ["/app/openrobo-api"]
