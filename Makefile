
BUILD_TIME := $(shell date "+%Y-%m-%d %H:%M:%S")
BUILD_VERSION := $(shell git rev-list --tags --max-count=1 | git describe --tags )

all:
	swag init
	go build -tags release -ldflags "-w -s \
		-X 'github.com/general252/goempty/pkg/version.BuildTime=${BUILD_TIME}' \
		-X 'github.com/general252/goempty/pkg/version.Version=${BUILD_VERSION}' " \
		-o bin/goempty.exe

lin:
	GOOS=linux GOARCH=amd64 go build -tags release -ldflags "-w -s \
		-X 'github.com/general252/goempty/pkg/version.BuildTime=${BUILD_TIME}' \
		-X 'github.com/general252/goempty/pkg/version.Version=${BUILD_VERSION}' " \
		-o bin/goempty
run:
	./bin/goempty.exe
