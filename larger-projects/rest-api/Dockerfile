# Stage 1: Build stage
FROM golang:alpine3.21 AS build

# Set Environment Variables
ENV CGO_ENABLED=1

# Set the working directory
WORKDIR /build

# Copy and download dependencies
RUN apk add --no-cache build-base
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o /app .

#################################

# Stage 2: Final stage
FROM alpine:3.21

# Copy the binary from the build stage
COPY --from=build /app /app

# CMD or ENTRYPOINT command
CMD ["/app"]