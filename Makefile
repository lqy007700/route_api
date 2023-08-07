
all: proto build docker docker-run

.PHONY: proto
proto:
	sudo docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) zxnl/protoc --proto_path=. --micro_out=. --go_out=:. ./proto/route_api/route_api.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/Cellar/go@1.19/1.19.11/bin/go build -o route_api *.go

.PHONY: docker
docker:
	sudo docker build . -t zxnl/route_api:latest

docker-run:
	sudo docker run -p 8085:8085 zxnl/route_api