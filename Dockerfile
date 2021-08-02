##Multi stage docker
FROM golang:1.16-alpine AS builder

ENV GO111MODULE=on

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

######## Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 