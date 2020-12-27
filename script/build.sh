#!/bin/sh

# build on Mac OS

go build -ldflags '-w -s' -o bin/lissajous ./cmd/lissajous

# regex lissajous version
ver=$(./bin/lissajous -v | grep -o "\d\.\d\.\d")

# package all executable files
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o bin/lissajous ./cmd/lissajous
zip 'tmp/lissajous'$ver'.darwin.zip' bin/lissajous

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o bin/lissajous ./cmd/lissajous
zip 'tmp/lissajous'$ver'.linux.zip' bin/lissajous

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o bin/lissajous.exe ./cmd/lissajous
zip 'tmp/lissajous'$ver'.windows.zip' bin/lissajous.exe
