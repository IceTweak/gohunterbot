# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY internal/ ./internal/
COPY cmd/ ./cmd/

# Build project
RUN go build -C ./cmd/gohunterbot/ -o ../../gohunterbot

# Bind to a TCP port
EXPOSE 4444

# Run
CMD [ "./gohunterbot" ]