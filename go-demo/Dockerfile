FROM golang:1.27.1@sha256:512690a5660563b57d37ecc31129e7f136e831db2aed24a1dbeb8ad7380dc0fa AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY main.go version.json ./

RUN mkdir -p build
RUN CGO_ENABLED=0 go build -v -o build ./...

# Declared as a named stage rather than referenced inline in COPY --from, so
# that Dependabot tracks it -- it only parses image references in FROM lines.
FROM busybox:1.38.0-uclibc AS busybox

FROM gcr.io/distroless/static:nonroot

COPY --from=busybox /bin/wget /usr/local/bin/wget
COPY --from=builder /app/build/go-demo /app/go-demo
COPY --from=builder /app/version.json /app/version.json

EXPOSE 8000

HEALTHCHECK --interval=3s --timeout=1s \
    CMD ["/usr/local/bin/wget", "--no-verbose", "--tries=1" , "--spider", "http://localhost:8000/__lbheartbeat__"]

CMD ["/app/go-demo"]
