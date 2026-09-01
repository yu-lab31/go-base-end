#!/bin/bash

set -e

MODULE_NAME=$(awk '/^module / {print $2}' go.mod)

if [ -z "$MODULE_NAME" ]; then
    echo "No module name found in go.mod"
    exit 1
fi

PROTO_FILES=$(find proto -name "*.proto")

if [ -z "$PROTO_FILES" ]; then
    echo "No .proto file found under proto folder"
    exit 0
fi

protoc -I=proto \
    --go_out=. --go_opt=module="${MODULE_NAME}" \
    --go-grpc_out=. --go-grpc_opt=module="${MODULE_NAME}" \
    ${PROTO_FILES}

