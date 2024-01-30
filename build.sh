#!/bin/sh

set -xe

go build -v -o ./.bin/qarwett ./cmd/bot/main.go

go build -v -o ./.bin/icalendar ./cmd/icalendar/main.go