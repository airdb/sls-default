SHELL = /bin/bash

VERSION:=$(shell git describe --dirty --always)
#VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse HEAD)
REPO := github.com/airdb/scf-go

LDFLAGS=-ldflags
LDFLAGS += "-X=$(REPO)/internal/version.Repo=$(REPO) \
            -X=$(REPO)/internal/version.Version=$(VERSION) \
            -X=$(REPO)/internal/version.Build=$(BUILD) \
            -X=$(REPO)/internal/version.BuildTime=$(shell date +%s)"

default: build deploy

build:
	GOOS=linux go build $(LDFLAGS) -o main main.go

deploy:
	swag init --generalInfo api/router.go --output ./docs
	SERVERLESS_PLATFORM_VENDOR=tencent GLOBAL_ACCELERATOR_NA=true sls deploy --stage test
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

release:
	swag init --generalInfo api/router.go --output ./docs
	sls deploy --stage release
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

log:
	sls logs --tail --stage test

logrelease:
	sls logs --tail --stage release

docker:
	docker run --platform linux/amd64 -it -v $(HOME)/go:/go airdb/serverless bash
