#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

export GOARCH="amd64"
export GOFLAGS="-mod=vendor"

FullPackage=$(go list -m)
Package=${FullPackage##*/}

for GOOS in darwin linux windows
do
    echo "Building $GOOS"
    export GOOS=$GOOS
    go build -o bin/${Package}.${GOOS}

done