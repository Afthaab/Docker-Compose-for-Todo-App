# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Run the Go app when the container starts
CMD ["./main"]
