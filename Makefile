VERSION = $(shell git describe --tags --abbrev=0 2>/dev/null)

all: build package

build:
	rm -rf toolkit-*/
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/lukecjohnson/toolkit/commands.version=$(VERSION) -s -w" -o toolkit-$(VERSION)/macos-amd64/toolkit
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X github.com/lukecjohnson/toolkit/commands.version=$(VERSION) -s -w" -o toolkit-$(VERSION)/macos-arm64/toolkit

package:
	rm -rf toolkit-*.zip
	zip -r -X toolkit-$(VERSION).zip toolkit-$(VERSION)
	rm -rf toolkit-*/