SHELL = /bin/bash

VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse HEAD)

LDFLAGS=-ldflags
LDFLAGS += "-X=airdb.io/airdb/scf-go/internal/version.Version=$(VERSION) \
            -X=airdb.io/airdb/scf-go/internal/version.Build=$(BUILD) \
            -X=airdb.io/airdb/scf-go/internal/version.BuildTime=$(shell date +%s)":

default: build deploy

build:
	GOOS=linux go build $(LDFLAGS) -o main main.go

deploy:
	sls deploy --stage test
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

release:
	sls deploy --stage release
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

docker:
	docker run -it --rm  --env-file ~/.env -v ~/:/go airdb/serverless bash
