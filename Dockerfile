FROM golang:1.25-alpine3.21 AS builder

RUN apk update && apk add --no-cache \
      ca-certificates \
      git \
      gcc \
      g++ \
      make \
      libc-dev \
      binutils \
      bash

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN cd internal/infrastructure/dds/dds && make -f Makefile_linux_static linux

ENV CGO_ENABLED=1
RUN go build -o bin/application ./cmd/app

FROM alpine:3.21 AS runner

RUN apk update && apk add --no-cache \
      ca-certificates \
      libstdc++ \
      libgcc \
      bash \
      && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY --from=builder /opt/bin/application ./

CMD ["./application"]
