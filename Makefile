
export GOOS=linux

all: headstart plugins

plugins: plugins/local.so plugins/aws.so

plugins/local.so:
	CGO_ENABLED=1 go build -buildmode=plugin -o local.so providers/local/local.go

plugins/aws.so:
	CGO_ENABLED=1 go build -buildmode=plugin -o aws.so providers/aws/aws.go

headstart:
	go build .

local: headstart
	HS_LOCAL_PATH=./sample_config.yml HS_PROVIDER_PATH=./ ./headstart
	
test:
	go test -v -covermode=count -coverprofile=coverage.out ./...

coverage-html: test
	go tool cover -html=coverage.out -o=coverage.html

.PHONY: headstart