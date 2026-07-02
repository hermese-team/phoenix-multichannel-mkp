FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o app ./marketplaces/shopee/cmd/consumer

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /src/app .
CMD ["./app"]
