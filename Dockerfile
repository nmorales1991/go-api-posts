# Start from the latest golang base image
FROM golang:buster

# Add Maintainer Info
LABEL maintainer="nmorales1991 <nmorales1991@example.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Tidy up the dependencies
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN go build -o main cmd/app/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
