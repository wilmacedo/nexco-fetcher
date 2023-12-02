filename=main

all: test build zip

build:
	@GOOS=linux GOARCH=amd64 go build -v -o bin/${filename} cmd/main.go

test:
	@go test ./...

zip:
	@rm -rf bin/${filename}.zip
	@zip -j bin/${filename}.zip bin/${filename}