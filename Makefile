SHELL = /bin/bash

default: build deploy

build:
	GOOS=linux go build -o main main.go

deploy:
	sls deploy --stage test
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

release:
	sls deploy --stage release
	@echo checkout all scf apps, https://serverless.cloud.tencent.com/

docker:
	docker run -it --rm  --env-file ~/.env -v $(shell pwd):/srv airdb/scf
