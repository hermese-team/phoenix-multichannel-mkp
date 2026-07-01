# ── Stage 1: builder ──────────────────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

ARG VERSION=dev
ARG COMMIT=none
ARG BUILD_TIME=unknown

WORKDIR /src

# Cache dependency download layer separately from source changes
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
      -ldflags="-s -w \
        -X github.com/ascend/phoenix-multichannel-mkp/pkg/build.Version=${VERSION} \
        -X github.com/ascend/phoenix-multichannel-mkp/pkg/build.Commit=${COMMIT} \
        -X github.com/ascend/phoenix-multichannel-mkp/pkg/build.BuildTime=${BUILD_TIME}" \
      -o /app/server \
      ./cmd/api

# ── Stage 2: runtime ──────────────────────────────────────────────────────────
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/server /server

EXPOSE 8080

ENV GIN_MODE=release

ENTRYPOINT ["/server"]
