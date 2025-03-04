#!/bin/sh

set -e

echo "Running tests via $0"

# setup the environment
export PATH=$PATH:$SNAP/go/bin
export CGO_ENABLED=0

echo "Copying project files to $SNAP_COMMON"
cp -r ./* $SNAP_COMMON

cd $SNAP_COMMON

echo "Running tests in $SNAP_COMMON"

go test -p 1 --cover "$@"
echo "✅ go test"

go vet "$@"
echo "✅ go vet"
