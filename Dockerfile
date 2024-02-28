# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY internal/ ./internal/
COPY cmd/ ./cmd/

# Build project
RUN go build -C ./cmd/gohunterbot/ -o ../../gohunterbot

FROM alpine

WORKDIR /build

# Copy binary
COPY --from=builder /build/gohunterbot /build/gohunterbot

# Bind to a TCP port
EXPOSE 4444

# Run
CMD [ "./gohunterbot" ]