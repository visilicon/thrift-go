export GOPATH:=$(CURDIR)/third:$(CURDIR)
export GOBIN:=$(CURDIR)/bin

all: install

fmt:
	gofmt -l -w -s src/

deps:
	go get git.apache.org/thrift.git/lib/go/thrift/...

install: deps
	go install server
	go install client

clean:
	rm -rf ./bin/server
	rm -rf ./bin/client
