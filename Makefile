run:
	./.bin/qarwett

brun: build
	./.bin/qarwett

test:
	go test -race ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

build:
	go build -v -o ./.bin/qarwett ./cmd/bot/main.go

docker-up: docker-build docker-compose.yml
	docker compose up -d

docker-up-dev: docker-compose.dev.yml
	docker compose -f docker-compose.dev.yml up -d

docker-build:
	docker build -t qarwett .

down:
	docker compose down