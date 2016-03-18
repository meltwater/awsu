export GO15VENDOREXPERIMENT=1

all: tools deps fmt build test lint

tools:
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/lint/golint
	go get -u github.com/Masterminds/glide

deps:
	glide install

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

build:
	 CGO_ENABLED=0 go build -o "awsu-`uname -s`-`uname -m`"
	 ln -sf "awsu-`uname -s`-`uname -m`" awsu

test:
	go test -bench=. -v -coverprofile=coverage.txt -covermode=atomic

lint:
	golint

clean:
	rm -f ./awsu

.PHONY: tools deps fmt build test lint clean
