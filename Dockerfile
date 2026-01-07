# Use Go 1.24 bookworm as build image
FROM golang:1.24-bookworm AS build

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY api ./

ENV GIN_MODE=release

RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -o dico

FROM gcr.io/distroless/base-debian11 AS release

COPY --from=build /build/dico /dico

# Document the port that may need to be published
EXPOSE 4242

USER nonroot
# Start the application
CMD ["/dico"]
