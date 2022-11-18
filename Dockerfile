FROM golang:alpine

RUN apk update --no-cache && \
    apk add tar --no-cache && \
    apk add curl --no-cache

ENV config=docker

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download && \
    go mod verify

COPY ./ ./

EXPOSE 8080

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz -C /bin

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT reflex -d none -s -R vendor. -r \.go$ -- go run cmd/main.go