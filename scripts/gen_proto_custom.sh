#!/bin/sh

set -eu

rm -rf custom/out
mkdir -p custom/out

protoc -I=custom \
  -I=$(go env GOPATH)/src \
  --go_out=custom/out \
  custom/custom.proto

cp custom/out/github.com/imura-tatsuya-ab/protoc-gen-option-example/custom/*.pb.go custom/
rm -rf custom/out
