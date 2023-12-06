all: build

build:
	mkdir bin
	go build -o bin/geassgo cmd/main.go