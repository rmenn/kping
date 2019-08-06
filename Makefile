SHELL:=/bin/bash
export PATH:=$(PATH):$(PWD)
export CGO_ENABLED:=0
export GOARCH:=amd64
export GOOS:=linux
export USER=rmenn
export REPOSITORY=kping
export VERSION=0.0.2
LDFLAGS=-X main.version=$(VERSION)

all: build docker
build:
	mkdir -p bin
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/kping

docker:
	docker build  . -t $(USER)/$(REPOSITORY):$(VERSION)
	docker push $(USER)/$(REPOSITORY):$(VERSION)
clean:
	rm -rf bin
