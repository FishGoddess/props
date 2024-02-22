.PHONY: test fmt

all: test

test:
	go test -cover -count=1 -test.cpu=1 ./...

fmt:
	go fmt ./...