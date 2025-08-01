# Use Go 1.24 bookworm as base image
FROM golang:1.24-bookworm AS base

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY api ./

ENV GIN_MODE=release

RUN go get .\
    & go build -o dico

# Document the port that may need to be published
EXPOSE 4242

# Start the application
CMD ["/build/dico"]
