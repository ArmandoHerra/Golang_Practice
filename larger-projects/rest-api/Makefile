#!/bin/bash -xe

# include .env
# export 

APP_NAME := go-api
DEV_DOCKERFILE := Dockerfile.dev
PROD_DOCKERFILE := Dockerfile
DEV_IMAGE := $(APP_NAME):dev
PROD_IMAGE := $(APP_NAME):latest
RUNNING_CONTAINER := $(shell docker ps --filter "name=$(APP_NAME)" --format "{{.ID}}" | head -n1)

# Build the production image
build-prod:
	docker build -f $(PROD_DOCKERFILE) -t $(PROD_IMAGE) .

# Build the development image
build-dev:
	docker build -f $(DEV_DOCKERFILE) -t $(DEV_IMAGE) .

# Run production container
run-prod:
	make build-prod && docker compose -f docker-compose.yml up

# Run development container
run-dev:
	make build-dev && docker compose -f docker-compose.dev.yml up

# Shell into running container
debug:
	docker exec -it $(RUNNING_CONTAINER) sh

# Clean all Docker cache and volumes
clean:
	docker system prune -a --volumes -f && docker builder prune -a -f

# Show status
status:
	docker ps --filter "name=$(APP_NAME)"
