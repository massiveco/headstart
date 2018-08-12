
export GOOS=linux

all: headstart plugins

plugins: plugins/local.so plugins/aws.so

plugins/local.so:
	CGO_ENABLED=1 go build -buildmode=plugin -o local.so providers/local/local.go

plugins/aws.so:
	CGO_ENABLED=1 go build -buildmode=plugin -o aws.so providers/aws/aws.go

headstart:
	go build .