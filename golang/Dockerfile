# Stage 1: Building the application and generating documentation
FROM golang:1.22-bullseye AS builder
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install swag CLI tool
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the entire source code into the image
COPY . .

# RUN ls -la

# Generate Swagger documentation
RUN /go/bin/swag init -g cmd/main.go 

# Build the Go application
RUN go build -o main cmd/main.go

# Stage 2: Setup the runtime container
FROM golang:1.22-bullseye
WORKDIR /app

# Copy the executable and Swagger docs from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Expose port (adjust if different)
# EXPOSE 8080

# Run the application
CMD ["./main"]
