APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=ghcr.io/mkdir28
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
TARGETARCH=$(shell uname -m)

ifeq ($(TARGETARCH),x86_64)
    override TARGETARCH=amd64
endif
ifeq ($(TARGETARCH),aarch64)
    override TARGETARCH=arm64
endif

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

deps:
	go get -v ./	

build: format
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X=github.com/mkdir28/kbot/cmd.appVersion=$(VERSION)"

linux:
	$(MAKE) build TARGETOS=linux TARGETARCH=amd64

arm:
	$(MAKE) build TARGETOS=linux TARGETARCH=arm64

windows:
	$(MAKE) build TARGETOS=windows TARGETARCH=amd64

macos:
	$(MAKE) build TARGETOS=darwin TARGETARCH=amd64

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH} --build-arg TARGETARCH=$(TARGETARCH)

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

clean:
	rm -rf kbot	
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}