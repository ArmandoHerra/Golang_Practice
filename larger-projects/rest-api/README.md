# How to run this project

## Prerequisites
- Docker
- Docker Compose

## Instructions

- Check the `Makefile` for some short commands used for developing/running the app.
- Use `$ make build-api` for building a container image with the API.
- Use `$ make run-api` for running the API with Docker Compose.
- Take note of the line in the `docker-compose.yaml` that uses the `Dockerfile.dev` file, which is intended to be used for development purposes only, we use `air-verse/air` for auto-recompilation of the code whenever we save code in our editors for quicker development, "production" `Dockerfile` uses multi-stage build to optimize the final image size and dependencies used.