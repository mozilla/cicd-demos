FROM golang:1.23@sha256:1acb493b9f9dfdfe705042ce09e8ded908ce4fb342405ecf3ca61ce7f3b168c7 AS builder

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
