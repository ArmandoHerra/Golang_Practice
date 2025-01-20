#!/bin/bash -xe

include .env
export 

RUNNING_CONTAINER=$(shell docker container list | sed -n '2 p' | awk '{print $$1}')

build-api:
	docker build . -t api:latest

run-api:
	docker compose up --build

debug-api:
	docker exec -it ${RUNNING_CONTAINER} sh

clean:
	docker image prune -af