FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o app ./sellchannel/shopee/cmd/server

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /src/app .
EXPOSE 8080
CMD ["./app"]
