# Use the official Go image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the project files into the container
COPY . .

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o main ./cmd

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["./main"]
