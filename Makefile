SHELL = /bin/bash
OUTPUT_DIR = objs
export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on
export GOSUMDB=off
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH=:.

SUB ?= clean build# and other plugins
# default APPVER or from $VAR or command-line
APPVER ?= 0.0.0

.PHONY: vendor ci build test


ci: $(SUB)
	@echo "$(SUB) build..."

vendor:
	go mod tidy && go mod vendor

build:
	go build -v -ldflags " \
		-X 'github.com/syzhang42/hermes/utils/ver.GitTag=$(shell git tag --sort=version:refname | tail -n 1)' \
		-X 'github.com/syzhang42/hermes/utils/ver.GitCommitLog=$(shell git log --pretty=oneline -n 1)' \
		-X 'github.com/syzhang42/hermes/utils/ver.BuildTime=$(shell date +'%Y.%m.%d.%H%M%S')' \
		-X 'github.com/syzhang42/hermes/utils/ver.Author=$(shell git log -1 --pretty=format:"%an)' \
		-X 'github.com/syzhang42/hermes/utils/ver.GoVersion=$(shell go version)' \
		-X 'github.com/syzhang42/hermes/utils/ver.Version=$(APPVER)' \
	" -o $(OUTPUT_DIR)/hermes main/main.go

test:
	@pushd $(OUTPUT_DIR) && \
		./hermes proxy -c ../conf/hermes.toml -v 1.0.0 && \
		popd

clean:
	rm -rf $(OUTPUT_DIR)
