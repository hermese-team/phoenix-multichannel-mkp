SHELL := /bin/bash

.PHONY: tidy build run-server run-consumer run-scheduler docker-build

tidy:
	go mod tidy

build:
	go build -o bin/shopee-server    ./sellchannel/shopee/cmd/server
	go build -o bin/shopee-consumer  ./sellchannel/shopee/cmd/consumer
	go build -o bin/shopee-scheduler ./sellchannel/shopee/cmd/scheduler

run-server:
	go run ./sellchannel/shopee/cmd/server

run-consumer:
	go run ./sellchannel/shopee/cmd/consumer

run-scheduler:
	go run ./sellchannel/shopee/cmd/scheduler

docker-build:
	docker build -f sellchannel/shopee/cmd/server/Dockerfile    -t shopee-server    .
	docker build -f sellchannel/shopee/cmd/consumer/Dockerfile  -t shopee-consumer  .
	docker build -f sellchannel/shopee/cmd/scheduler/Dockerfile -t shopee-scheduler .
