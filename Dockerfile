ARG GOVERSION=1.25.5

FROM golang:${GOVERSION} AS dev
RUN go install "github.com/air-verse/air@latest" && \
    go install "github.com/pressly/goose/v3/cmd/goose@latest"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
VOLUME /storage
ARG DEBUG
ENV DEBUG=${DEBUG:-}
ENV LISTEN_ADDR=":1991"
ENV DATABASE_PATH="/storage/openswingmusic.db"
ENV CGO_ENABLED=1
ARG SWINGSONIC_BASE_URL
ENV SWINGSONIC_BASE_URL=${SWINGSONIC_BASE_URL:-}
ENV CREDENTIALS_PROVIDER="file"
ENV USERS_FILE_PATH="users"
CMD ["air", "-c", ".air.toml"]
