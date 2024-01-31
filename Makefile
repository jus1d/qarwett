run:
	./.bin/qarwett

brun: build
	./.bin/qarwett

build:
	go build -v -o ./.bin/qarwett ./cmd/bot/main.go
	go build -v -o ./.bin/icalendar ./cmd/icalendar/main.go

docker-up: docker-build docker-compose.yml
	docker compose up -d

docker-up-dev: docker-build docker-compose.dev.yml
	docker compose -f docker-compose.dev.yml up -d

docker-build:
	docker build -t qarwett .

down:
	docker compose down