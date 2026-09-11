FROM golang:1.27.1@sha256:f44f6e88636cfb311f9ebace870ded69d943f227bb3cb27d32ffd84ea18c43ea AS builder

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
