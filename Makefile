GOPATH := $(shell pwd)
PKG_NAME := coconut.com

all:
	@mkdir -p bin/  	
	go build -o bin/deploygate $(PKG_NAME) 

test:
	go test $(PKG_NAME)/...

