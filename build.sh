#!/bin/sh

set -xe

go mod download

go build -v -o ./.bin/qarwett ./cmd/bot/main.go