FROM golang:1.23.7@sha256:cb45cf739cf6bc9eaeacf75d3cd7c157e7d39b757216d813d8115d026ee32e75 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN mkdir -p build
RUN CGO_ENABLED=0 go build -v -o build ./...

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/build/go-demo /app/go-demo
COPY --from=builder /app/version.json /app/version.json

EXPOSE 8000

CMD ["/app/go-demo"]
