FROM golang:1.20.5 AS builder

COPY . /src
WORKDIR /src

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/ipam-api .

# Pass in name as --build-arg
FROM gcr.io/distroless/static:nonroot
# `nonroot` coming from distroless
USER 65532:65532

COPY  --from=builder /src/bin/ipam-api /ipam-api

# Run the web service on container startup.
ENTRYPOINT ["/ipam-api"]
CMD ["serve"]