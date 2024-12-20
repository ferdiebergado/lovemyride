# Use Go 1.22 bookworm as base image
FROM golang:1.22-bookworm AS base

# Development stage
# =============================================================================
# Create a development stage based on the "base" image
FROM base AS development

# Change the working directory to /app
WORKDIR /app

# Install the air CLI for auto-reloading
RUN go install github.com/air-verse/air@v1.52.2

# Copy the go.mod and go.sum files to the /app directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Start air for live reloading
CMD ["air"]

# Testing stage
# =============================================================================
# Create a testing stage based on the "base" image
FROM base AS testing

# Change the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the /app directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Start the test
RUN go test -v -race ./...

# Builder stage
# =============================================================================
# Create a builder stage based on the "base" image
FROM base AS builder

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -o lovemyride

# Production stage
# =============================================================================
# Create a production stage to run the application binary
FROM scratch AS production

# Move to working directory /prod
WORKDIR /prod

# Copy binary from builder stage
COPY --from=builder /build/lovemyride ./

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["/prod/lovemyride"]
