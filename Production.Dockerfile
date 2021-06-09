FROM golang:1.16-buster as builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o cmd/korai main.go

FROM debian:10-slim
WORKDIR /app/
COPY --from=builder /build/cmd/korai .
EXPOSE 9000
CMD ["/app/korai"]
