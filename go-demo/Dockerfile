FROM golang:1.23.12@sha256:60deed95d3888cc5e4d9ff8a10c54e5edc008c6ae3fba6187be6fb592e19e8c0 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY main.go version.json ./

RUN mkdir -p build
RUN CGO_ENABLED=0 go build -v -o build ./...

FROM gcr.io/distroless/static:nonroot

COPY --from=busybox:1.37.0-uclibc /bin/wget /usr/local/bin/wget
COPY --from=builder /app/build/go-demo /app/go-demo
COPY --from=builder /app/version.json /app/version.json

EXPOSE 8000

HEALTHCHECK --interval=3s --timeout=1s \
    CMD ["/usr/local/bin/wget", "--no-verbose", "--tries=1" , "--spider", "http://localhost:8000/__lbheartbeat__"]

CMD ["/app/go-demo"]
