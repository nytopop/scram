default: build

build:
	go fmt *.go
	go build -o bin/scram

install:
	go fmt *.go
	go install

run:
	go fmt *.go
	go build -o bin/scram
	bin/scram
