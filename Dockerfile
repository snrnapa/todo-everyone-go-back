# Stage 1: Build the Go application
FROM golang:1.21.4-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code to the container
COPY . .

# Build the Go application
RUN go build -o /todo-everyone ./cmd/api

# Stage 2: Create a smaller image to run the Go app
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go application from the builder stage
COPY --from=builder /todo-everyone .

# Copy the credentials file
COPY serviceAccountKey.json .

# Expose the application on port 8080
EXPOSE 8080

# Command to run the application
CMD ["./todo-everyone"]