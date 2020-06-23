FROM golang:1.14.1-alpine3.11 as builder
LABEL maintainer="Agostin Jean-baptiste <Jbagostin@gmail.com>"

ENV GO111MODULE=on

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api .

EXPOSE 8080

ENTRYPOINT ["./app"]