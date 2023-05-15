build:
	go build -o sent

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sent

.PHONY: build build-linux