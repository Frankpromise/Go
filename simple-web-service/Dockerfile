FROM golang:1.22-alpine3.19
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod /app/
RUN go mod download

# Copy the source code
COPY *.go /app/

# Build the Go application
RUN go build -o /demo

# Expose the port the application runs on
EXPOSE 3000

# Command to run the executable
CMD [ "/demo" ]
