# Use the official Golang image with version 1.21
FROM golang:1.21

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 9090

# Command to run the application
CMD ["./main"]
