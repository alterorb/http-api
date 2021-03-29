FROM golang:1.16-alpine AS builder
WORKDIR /build
COPY *.go go.mod go.sum /build/
RUN GOOS=linux GOARCH=amd64 go build -o http-api

FROM alpine
WORKDIR /app
COPY --from=builder /build/http-api .
COPY openapi openapi
EXPOSE 8080
CMD ["./http-api"]