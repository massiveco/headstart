export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64

all: headstart plugins

plugins: plugins/local.so plugins/aws.so

plugins/local.so:
	cd providers/local && go build -buildmode=plugin -o local.so

plugins/aws.so:
	cd providers/aws && go build -buildmode=plugin -o aws.so

headstart:
	go build .