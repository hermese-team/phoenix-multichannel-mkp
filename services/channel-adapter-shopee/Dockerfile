# ==========================================
# STEP 1: Base & Dependencies
# ==========================================
FROM golang:1.25-alpine AS base

RUN apk add --no-cache make curl git

# Private module access
RUN go env -w GOPRIVATE=git.amaze-x.com/*
ARG USER
ARG ACCESS_TOKEN
RUN git config --global url."https://${USER}:${ACCESS_TOKEN}@git.amaze-x.com".insteadOf "https://git.amaze-x.com"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# ==========================================
# STEP 2: Builder
# ==========================================
FROM base AS builder
COPY . .

RUN go build -o /app/build/server ./cmd/server
RUN go build -o /app/build/worker ./cmd/worker

# ==========================================
# STEP 3: Runtime image
# ==========================================
FROM alpine:latest

RUN apk add --no-cache tzdata ca-certificates

WORKDIR /app
COPY --from=builder /app/build /app

# Seed config structure — values are overridden at runtime by ExternalSecret env vars.
COPY config.example.json /app/config.json

# EXEC is set by k8s Deployment to "server" or "worker"
CMD ["sh", "-c", "/app/${EXEC}"]
