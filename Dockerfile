ARG GO_VERSION=1.25.6
ARG ALPINE_VERSION=3.23


FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS dev
# mime type support
RUN apk update && apk add mailcap && rm -rf /var/cache/apk/*
VOLUME /storage
ENV LISTEN_ADDR=":1991"
ENV DATABASE_PATH="/storage/openswingmusic.db"
ENV CRED_PROVIDER="database"
RUN go install "github.com/air-verse/air@latest"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
CMD ["air", "-c", ".air.toml"]


FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder
ARG APP_VERSION
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build \
        -ldflags "-X github.com/tikhonp/openswingsonic/internal/util.AppVersion=${APP_VERSION}" \
        -o /bin/oswingsonic main.go


FROM alpine:${ALPINE_VERSION} AS prod

# mime type support
RUN apk update && apk add mailcap && rm -rf /var/cache/apk/*

ARG VERSION
ARG BUILD_DATE
ARG VCS_REF
ARG APP_VERSION

LABEL org.opencontainers.image.title="OpenSwingMusic" \
      org.opencontainers.image.description="A translation layer allowing clients compatible with Open Subsonic API to work with Swing Music" \
      org.opencontainers.image.version=${APP_VERSION} \
      org.opencontainers.image.source="https://github.com/tikhonp/openswingsonic" \
      org.opencontainers.image.licenses="AGPL-3.0" \
      org.opencontainers.image.created=${BUILD_DATE} \
      org.opencontainers.image.revision=${VCS_REF}

VOLUME /storage
ENV LISTEN_ADDR=":1991"
ENV DATABASE_PATH="/storage/openswingmusic.db"
ENV CRED_PROVIDER="env"

EXPOSE 1991
WORKDIR /app

COPY --from=builder /bin/oswingsonic /bin/oswingsonic

CMD ["/bin/oswingsonic"]
