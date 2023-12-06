GOOS ?= linux
GOARCH ?= amd64
OUTPUT ?= geassgo-${GOOS}-${GOARCH}
GO111MODULE ?= on

all: build

clean:
	rm -rf bin/${OUTPUT}
	rm -rf bin/${OUTPUT}.tar.gz

dep:
	go mod download

build: clean dep
	mkdir -p bin/${OUTPUT}
	GOOS=${GOOS} GOARCH=${GOARCH} GO111MODULE=${GO111MODULE} go build -v -o bin/${OUTPUT}/geassgo cmd/main.go
	cd bin && tar -zcvf ${OUTPUT}.tar.gz ${OUTPUT}