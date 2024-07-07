# Use the latest Go image
FROM golang:latest AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Install necessary runtime dependencies in Alpine
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Install Go runtime and other necessary dependencies
RUN apk add --no-cache go ca-certificates

# Copy the source code from the build stage
COPY --from=build /app .

# Command to run the Go application
CMD ["go", "run", "main.go"]
