#!/bin/sh

#
# Copyright 2020 Singularity, Inc. All rights reserved.
#

# build on Mac OS

# build darwin OS execable file
go build -ldflags "-s -w" -o lissajous .

# regex lissajous version
ver=$(./lissajous -v | grep -o "\d\.\d\.\d")

# package darwin execable file
zip 'lissajous'"$ver"'.darwin.zip' lissajous

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o lissajous.exe .
zip 'lissajous'"$ver"'.windows.zip' lissajous.exe

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o lissajous .
zip 'lissajous'"$ver"'.linux.zip' lissajous
