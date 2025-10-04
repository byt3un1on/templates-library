
install:
	go clean -cache
	go mod tidy

build:
	go build -o btu

test:
	go test

cover:
	go test -coverpkg=./... -coverprofile=coverage.out;
	go tool cover -html=coverage.out -o coverage.html;

run:
	go run .

fmt:
	go fmt ./...

vet:
	go vet ./...

deploy: build
	mv btu /usr/local/go/bin/

clean:
	rm btu || true
	rm coverage.out || true
	rm coverage.html || true
	go clean -cache
