build:
	go build -o send 

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o send 

.PHONY: build build-linux
