#-----------------------------------------------------------------------------
# Global Variables
#-----------------------------------------------------------------------------
SHELL=bash

APP_VERSION := latest

#-----------------------------------------------------------------------------
# BUILD
#-----------------------------------------------------------------------------

.PHONY: default build test publish build_local lint
default: depend test lint build 

depend:
	go get -u github.com/golang/dep
	dep ensure
test:
	go test -v ./...
build:
	go build 
lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter ./... --exclude=vendor --deadline=60s


#-----------------------------------------------------------------------------
# CLEAN
#-----------------------------------------------------------------------------

.PHONY: clean 

clean:
	rm -rf nats-consumer

