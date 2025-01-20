# Build Stage
FROM golang:alpine3.21 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache build-base

COPY . .
RUN CGO_ENABLED=1 go build -o /app .

# Final Stage
FROM alpine:3.21

COPY --from=builder /app /app
CMD ["/app"]
