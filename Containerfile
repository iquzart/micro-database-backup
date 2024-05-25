# Stage 1: Build the Go application
FROM golang:1.21 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o micro-db-backup cmd/main.go

# Stage 2: Create the final image with Red Hat UBI 9 and mongodump
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

# Install mongodump
RUN microdnf install -y mongodb-database-tools && microdnf clean all

# Copy the binary from the builder stage
COPY --from=builder /app/micro-db-backup /usr/local/bin/micro-db-backup

# Set the entry point for the container
ENTRYPOINT ["/usr/local/bin/micro-db-backup"]

