# Stage 1: Build the Go application
FROM golang:1.22.5 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o knapsack

# Stage 2: Create a minimal image with the binary
FROM alpine:3.18

# Set working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/knapsack .

# Make the binary executable
RUN chmod +x knapsack

# Expose a port if necessary (Optional: Only if you plan to serve the app via a web server)
EXPOSE 8080

# Command to run the Go program
CMD ["./knapsack"]
