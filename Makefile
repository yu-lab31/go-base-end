.PHONY: default build test clean init

default: build

build:
	mkdir build

test:
	go clean --testcache
	go test ./test/logger/... -v

clean:
	go clean --testcache
	$(RM) -r build

init:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make init name=<module name>"; \
		exit 1; \
	fi
	@./hack/init.sh $(name)
