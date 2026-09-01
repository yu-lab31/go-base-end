.PHONY: default build test clean init pb

default: build

build: pb
	mkdir build

pb:
	@./hack/protoc.sh

test:
	go clean --testcache
	go test ./test/logger/... -v
	go test ./test/resource/... -v

clean: clean_pb
	go clean --testcache
	$(RM) -r build

clean_pb:
	find . -type f \( -name "*.pb.go" -o -name "*_grpc.pb.go" \) -delete

init:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make init name=<module name>"; \
		exit 1; \
	fi
	@./hack/init.sh $(name)
	@./hack/protoc.sh
