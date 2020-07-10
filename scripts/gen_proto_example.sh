#!/bin/sh

set -eu

rm -rf proto/out
mkdir -p proto/out

protoc -I=proto \
  -I=. \
  -I=$(go env GOPATH)/src \
  --go_out=proto/out \
  --custom_out=. \
  proto/example.proto

cp proto/out/github.com/imura-tatsuya-ab/protoc-gen-option-example/proto/*.pb.go proto/
rm -rf proto/out
