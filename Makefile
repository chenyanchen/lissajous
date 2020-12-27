# Directory
BINDIR = $(CURDIR)/bin
TMPDIR = $(CURDIR)/tmp

# applications
LISSAJOUS = lissajous
DIRECTORY_CLEANER = dir_cleaner

# go options
GOFLAGS   =
GOLDFLAGS = -w -s

#  build
.PHONY: build
build: build-lissajous build-dir-cleaner

# TODO: build executable file with scripts
build-lissajous:
	go build $(GOFLAGS) -ldflags '$(GOLDFLAGS)' -o '$(BINDIR)'/$(LISSAJOUS) ./cmd/lissajous

build-dir-cleaner:
	go build $(GOFLAGS) -ldflags '$(GOLDFLAGS)' -o '$(BINDIR)'/$(DIRECTORY_CLEANER) ./cmd/directory_cleaner

.PHONY: release
release:
	./script/build.sh
