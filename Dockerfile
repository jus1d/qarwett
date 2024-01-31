FROM golang:1.21.5-alpine3.18 AS builder

RUN go version

COPY ./ /qarwett
WORKDIR /qarwett

RUN go mod download
RUN ./build.sh

# Lightweight docker container with binary files
FROM alpine:latest

WORKDIR /app

COPY --from=builder /qarwett/.bin/ ./bin
COPY --from=builder /qarwett/config/ ./config

CMD ["./bin/qarwett"]