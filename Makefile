.PHONY: install build build-all linux-amd64

install:
	go mod download

build:
	go build -ldflags "-s -w" -o build/osm-bot cmd/main.go

build-all: linux-amd64

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o build/osmbot-linux-amd64 cmd/main.go
