
all: headstart plugins

plugins: plugins/local.so plugins/aws.so

plugins/local.so:
	cd providers/local && CGO_ENABLED=1 go build -buildmode=plugin -o ../local.so

plugins/aws.so:
	cd providers/aws && CGO_ENABLED=1 go build -buildmode=plugin -o ../aws.so

headstart:
	CGO_ENABLED=1 go build .