build:
	go build -o query

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o query 

.PHONY: build build-linux
