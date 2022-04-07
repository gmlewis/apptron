ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
GO_FILES := $(shell find . -name "*.go") 

apptron: clientjs/dist/client.js $(GO_FILES)
	go build -o ./apptron ./cmd/apptron/main.go

debug-pkg: $(GO_FILES)
	go build -tags pkg -o ./debug-pkg ./cmd/debug

debug-app: clientjs/dist/client.js $(GO_FILES) 
	go build -tags app -o ./debug-app ./cmd/debug

debug-cmd: apptron $(GO_FILES)
	go build -tags cmd -o ./debug-cmd ./cmd/debug

clientjs/dist/client.js: clientjs/lib/*.js clientjs/src/*.ts
	make -C clientjs build

.PHONY: clean
clean:
	rm -rf ./debug-pkg ./debug-app ./debug-cmd ./apptron
