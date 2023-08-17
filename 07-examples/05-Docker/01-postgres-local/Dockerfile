# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local Go application files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Run the Go application
CMD ["./main"]
