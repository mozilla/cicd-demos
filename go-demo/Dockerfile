FROM golang:1.23.10@sha256:7d3f5bf3aee7f83298bda8615d6914916c4b5408819557ea3f54404d9b5b9798 AS builder

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
