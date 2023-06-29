FROM golang:1.18-alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o cmd/httserver main.go

FROM alpine
WORKDIR /
COPY --from=builder /build/cmd .
RUN chmod +x httserver
EXPOSE 8999
ENTRYPOINT ["./httserver"]
