APP_NAME=mini-ticket

run:
	go run main.go

build:
	go build -o bin/$(APP_NAME)

swagger:
	swag init

watch:
	air

test:
	go test ./...

deps:
	go mod tidy

lint:
	golangci-lint run

dev:
	make swagger
	make watch