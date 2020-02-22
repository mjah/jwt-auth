# ---
# Compile source
# ---
FROM golang:1.13-alpine AS builder

# Set working directory within $GOPATH/src/
WORKDIR /go/src/jwt-auth

# Copy source over
COPY . .

# Install build dependencies
RUN apk add --no-cache git build-base

# Get missing packages
RUN go get -d -v ./...

# Run tests
RUN go test ./...

# Compile source and install binary to $GOPATH/bin/
RUN go install -v ./...

# ---
# Run executable
# ---
FROM alpine:latest

# Copy binary from previous stage to root
COPY --from=builder /go/bin/jwt-auth /jwt-auth

# Set entry point to executable
ENTRYPOINT [ "/jwt-auth" ]

# Expose port required by executable
# - commented out, expose via argument instead
# EXPOSE 9096

# Maintainer
LABEL maintainer "Muhammad Abdul Hafiz <moe@moejay.com>"
