FROM gcr.io/distroless/static

# Copy the binary that goreleaser built
COPY  ipam-api /ipam-api

# Run the web service on container startup.
ENTRYPOINT ["/ipam-api"]
CMD ["serve"]
