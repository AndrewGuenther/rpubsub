all: install

deps:
	go get ./...

build: deps
	go build ./...

install: deps
	go install ./...

clean:
	go clean ./...

