
install:
	go clean -cache
	go mod tidy

build:
	go build -o btu

test:
	go test

run:
	go run .
