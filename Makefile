LDFLAGS += -X "fingerprintClient/config.BuildTime=$(shell date "+%Y-%m-%d %T %Z")"
LDFLAGS += -X "fingerprintClient/config.GitCommit=$(shell git rev-parse HEAD)"
LDFLAGS += -X "fingerprintClient/config.GitStatus=$(shell if git status | grep -q 'clean'; then echo clean; else echo dirty; fi)"
.PHONY: build clean

all: build

build:
	@go build -ldflags '$(LDFLAGS)' -o ./bin/fp_client ./main/server.go

clean:
	@rm -rf ./bin/*
