FROM golang:1.25-alpine3.21 AS builder

RUN apk update && apk add --no-cache \
      ca-certificates \
      git \
      gcc \
      g++ \
      libc-dev \
      binutils \
      bash

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o bin/application ./cmd/app

FROM alpine:3.21 AS runner

RUN apk update && apk add --no-cache ca-certificates libc6-compat openssh bash && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY --from=builder /opt/bin/application ./
COPY --from=builder /opt/internal/config/config_local.yaml ./

CMD ["./application"]
