ARG GO_VERSION=1.25.5
ARG ALPINE_VERSION=3.23


FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS dev
RUN apk add --no-cache build-base
VOLUME /storage
ENV LISTEN_ADDR=":1991"
ENV DATABASE_PATH="/storage/openswingmusic.db"
ENV CREDENTIALS_PROVIDER="database"
ENV CGO_ENABLED=1
RUN go install "github.com/air-verse/air@latest" && \
    go install "github.com/pressly/goose/v3/cmd/goose@latest"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
CMD ["air", "-c", ".air.toml"]


FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder
ARG APP_VERSION
RUN apk add --no-cache build-base
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=1 \
    go build \
        -ldflags "-X github.com/tikhonp/openswingsonic/internal/util.AppVersion=${APP_VERSION}" \
        -o /bin/oswingsonic main.go


FROM alpine:${ALPINE_VERSION} AS prod

ARG VERSION
ARG BUILD_DATE
ARG VCS_REF
ARG APP_VERSION

LABEL org.opencontainers.image.title="OpenSwingMusic"
LABEL org.opencontainers.image.description="A translation layer allowing clients compatible with Open Subsonic API to work with Swing Music"
LABEL org.opencontainers.image.version=${APP_VERSION}
LABEL org.opencontainers.image.source="https://github.com/tikhonp/openswingsonic"
LABEL org.opencontainers.image.licenses="AGPL-3.0"
LABEL org.opencontainers.image.created=${BUILD_DATE}
LABEL org.opencontainers.image.revision=${VCS_REF}

VOLUME /storage
ENV LISTEN_ADDR=":1991"
ENV DATABASE_PATH="/storage/openswingmusic.db"
ENV CREDENTIALS_PROVIDER="database"

EXPOSE 1991
WORKDIR /app

COPY --from=builder /bin/oswingsonic /bin/oswingsonic

CMD ["/bin/oswingsonic"]
