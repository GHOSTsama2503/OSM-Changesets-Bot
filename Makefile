install:
	go mod tidy

build-all: linux-amd64

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o build/osmbot-linux-amd64 cmd/main.go
