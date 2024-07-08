HOMEDIR:=$(shell pwd)
OUTDIR:=$(HOMEDIR)


export GOENV = $(HOMEDIR)/go.env


all: prepare complie package

prepare:
	git version
	go env
	go mod download || go mod download -x

complie:build

build:prepare
	go build -o $(HOMEDIR)/mini-spider main.go

package:
	mkdir -p $(OUTDIR)/bin
	mv mini-spider $(OUTDIR)/bin/mini-spider
	mkdir $(OUTDIR)/output
	chmod 777 output
test: prepare
    go test -race -timeout=120s -v -cover $(GOPKGS) -coverprofile=coverage.out | tee unittest.txt
clean:
	rm -rf $(OUTDIR)/bin
	rm -rf $(OUTDIR)/output

.PHONY: all prepare compile test package  clean build test