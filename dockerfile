# Start from the official Golang image to create a build artifact.
FROM golang:1.20

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Expose port 8080 to the outside world.
EXPOSE 8080

# cd to cmd/api
WORKDIR /app/cmd/api


# Command to run go program.
CMD ["go", "run", "main.go"]
