FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o app ./sellchannel/shopee/cmd/consumer

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /src/app .
CMD ["./app"]
