# Stage 1: Build stage
FROM golang:1.22-bullseye AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN apt-get update && apt-get install -y build-essential && go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 go build -o /bin/myapi .

#################################

# Stage 2: Final stage
FROM gcr.io/distroless/base

# Copy the binary from the build stage
COPY --from=build /bin/myapi /app

# CMD or ENTRYPOINT command
CMD ["/app"]