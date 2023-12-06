all: build

clean:
	rm -rf bin

dep:
	go mod download

build: clean dep
	mkdir -p bin
	go build -v -o bin/geassgo cmd/main.go