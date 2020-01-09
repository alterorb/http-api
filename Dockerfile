FROM golang:1.13.5 AS builder
WORKDIR /build
COPY *.go go.mod go.sum /build/
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -linkmode external -extldflags -static" -o http-api

FROM scratch
WORKDIR /app
COPY --from=builder /build/http-api .
COPY openapi openapi
EXPOSE 8080
CMD ["./http-api"]