# Use golang 1.18 on Alpine Linux as the base image for the build stage
FROM golang:1.19.8-alpine AS builder

# Install git, necessary for fetching Go dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files into the container
COPY go.mod .
COPY go.sum .

# Download the Go modules specified in go.mod and go.sum
RUN go mod download

# Copy all Go source files from the current directory into the container
COPY *.go ./

# Build the application: disable CGO, set the target OS to Linux, compile for all dependencies, and output the binary named 'main'
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use the latest version of Alpine Linux for the final image
FROM alpine:latest

# Add a new group 'appgroup' and user 'gouser' for running the application
RUN addgroup -S appgroup && adduser -S gouser -G appgroup

# Set the working directory in the container to /app
WORKDIR /app

# Copy the compiled binary from the builder stage into the current container
COPY --from=builder /app/main .

# Switch to the user 'gouser' for security purposes (non-root user)
USER gouser

# Inform Docker that the container listens on the specified network port at runtime
EXPOSE 8080

# Define the command to run the application
CMD ["./main"]
