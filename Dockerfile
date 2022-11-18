FROM golang:alpine

RUN apk update --no-cache

ENV config=docker

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download && \
    go mod verify

COPY ./ ./

EXPOSE 8080

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT reflex -d none -s -R vendor. -r \.go$ -- go run cmd/main.go