ARG GOVERSION=1.25.5


FROM golang:${GOVERSION} AS base
VOLUME /storage
ENV LISTEN_ADDR=":1991"
ENV DATABASE_PATH="/storage/openswingmusic.db"
ENV CREDENTIALS_PROVIDER="database"
ENV CGO_ENABLED=1


FROM base AS dev
RUN go install "github.com/air-verse/air@latest" && \
    go install "github.com/pressly/goose/v3/cmd/goose@latest"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
CMD ["air", "-c", ".air.toml"]
