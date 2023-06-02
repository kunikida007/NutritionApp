FROM golang:1.20-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV GO111MODULE=on
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN apk update && apk add curl

RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080


CMD ["air", "-c", ".air.toml"]

