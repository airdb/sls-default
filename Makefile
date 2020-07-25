all: deploy

deploy:
	GOOS=linux go build -o main main.go
	scf deploy -t template.yaml  -f
