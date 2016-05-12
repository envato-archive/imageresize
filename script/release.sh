#!/bin/bash

version=${1:-$(git describe --tags --candidates=1 --dirty)}

if [ -z "$version" ]; then
    echo "Usage: $0 [version]" >&2
    exit 1
fi

dir="pkg/${version}"
prefix="${dir}/imageresize-${version}"

build() {
    local os=$1
    local arch=$2
    local extension=$([ "xwindows" = "x${os}" ] && echo ".exe")
    local binary=${prefix}-${os}-${arch}${extension}

    echo "Building $binary"
    GOOS=$os GOARCH=$arch go build -o $binary .
}

mkdir -p $dir

build windows 386
build windows amd64
build linux 386
build linux amd64
build darwin 386
build darwin amd64
