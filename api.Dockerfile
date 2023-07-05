# syntax=docker/dockerfile:1

FROM golang:alpine AS build-stage

WORKDIR /images

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o . ./cmd/api/main

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /images/. /app/

CMD ["./main"]
