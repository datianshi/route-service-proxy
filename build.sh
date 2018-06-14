#!/bin/bash

set -e

PROJECT_DIR="$(go env GOPATH)/src/github.com/datianshi/route-service-proxy"
BUILD_DIR="$PROJECT_DIR/build"

mkdir -p $BUILD_DIR

pushd $PROJECT_DIR
  go test ./...
popd

go test ./...
pushd $PROJECT_DIR/cmd/
  GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/routing-linux
#  GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/routing-mac
#  GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/routing-win64
popd
