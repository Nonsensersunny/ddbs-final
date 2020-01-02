PROJECT=ddbs-final
VERSION=1.0

proxy:
	export GOPROXY=https://goproxy.io

pre:
	export GOPATH=${PWD}
	export GO111MODULE=on
	go mod vendor

build-travis: pre
	go build cmd/main.go
	docker-compose --build -d

build-jenkins: proxy pre
	go build cmd/main.go
	docker-compose --build -d

build-local:
	go build cmd/main.go