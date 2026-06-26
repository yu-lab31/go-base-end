.PHONY: default build test clean

default: build

build:
	mkdir build

test:
	go clean --testcache
	go test ./test/logger/... -v

clean:
	go clean --testcache
	$(RM) -r build
