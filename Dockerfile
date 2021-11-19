FROM golang:1.16.5-alpine3.13 AS build_base

RUN apk add --no-cache ca-certificates curl git openssh build-base

WORKDIR /bakery

COPY go.mod .
COPY go.sum .

FROM build_base AS builder

COPY . .
RUN go build -mod=vendor -ldflags "-X github.com/cbsinteractive/bakery/handlers.GitSHA=$(git rev-parse HEAD)" -o bakery cmd/http/main.go

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app/

RUN apk add --no-cache tzdata

COPY --from=builder bakery .

RUN adduser -D bakery
USER bakery

EXPOSE 8080

ENTRYPOINT ["./bakery"]
