.ONESHELL:
ifeq ($(OS),Windows_NT)
SHELL=C:/Program Files/Git/bin/bash.exe
build: windows
else
SHELL=/usr/bin/bash
build: linux
endif

windows: prepare
	export GOOS=windows
	export GOARCH=amd64
	export CGO_ENABLED=1
	go build -trimpath --tags "osusergo,netgo,sqlite_omit_load_extension" -ldflags "-s -w" -o ./dist/passport.exe .

linux: prepare
	export GOOS=linux
	export GOARCH=amd64
	export CGO_ENABLED=1
	go build -trimpath -a --tags "osusergo,netgo,sqlite_omit_load_extension" -ldflags "-s -w" -o ./dist/passport .

prepare:
	@rm -rf ./dist/
	@mkdir -p ./dist/
	@cp -r ./configs ./dist/
	@cp -r ./web ./dist/

.PHONY: build windows linux
.DEFAULT_GOAL=build
