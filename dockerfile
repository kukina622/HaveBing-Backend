FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN  go mod download

COPY . .
RUN go test ./... && go build -o /bin/server cmd/server.go

FROM alpine:latest AS release
COPY --from=builder /bin/server /bin/server
ENTRYPOINT [ "/bin/server" ]