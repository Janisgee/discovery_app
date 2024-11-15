# What base image to use for this application
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY src/go.mod ./
RUN go mod download  
# Skip if no dependencies are specified in go.mod

# Copy the rest of the application code into the image
COPY src/ ./src

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/discovery_app ./src/main.go

# To bind to a TCP port
EXPOSE 8080

# Command when the container runs
CMD ["/app/discovery_app"]


