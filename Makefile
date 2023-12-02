filename=fetcher

all: test build

build:
	@go build -v -o bin/${filename} cmd/main.go

test:
	@go test ./...