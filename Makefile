GO=go
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))
GOFLAGS= 
BIN=bin/mesos-to-cloudwatch
VERSION := $(shell git rev-parse HEAD)
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
IMAGE="axelspringer/mesos-to-cloudwatch"

all: build/local

help:
	@echo 'Available commands:'
	@echo
	@echo 'Usage:'
	@echo '    make restore     	Install go deps.'
	@echo '    make build/local    	Compile the project.'
	@echo '    make test/local    	Run ginkgo test suites.'
	@echo '    make build/docker    Create docker container'
	@echo '    make clean    		Clean the directory tree.'
	@echo

build/local: 
	$(GO) get ./...
	$(GO) build -ldflags "-X main.Version=$(VERSION)" -o $(BIN) $(GOFLAGS) $(RACE) main.go

deploy/local: build/linux
	docker restart `docker ps | grep "/mesos-to-cloudwatch" | awk '{printf $$1}'`

build/linux:
	$(GO) get ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags "-X main.Version=$(VERSION)" -o "$(BIN)_linux" $(GOFLAGS)  main.go

build/docker: build/linux
	docker build -t $(IMAGE) .

compose/up: build/linux
	docker-compose -f docker-compose.yml up -d

compose/down:
	docker-compose -f docker-compose.yml down

push/docker:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push $(IMAGE)

restore:
	dep ensure
