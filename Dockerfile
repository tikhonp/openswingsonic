ARG GOVERSION=1.25.5

FROM golang:${GOVERSION}-alpine AS dev
RUN go install "github.com/air-verse/air@latest"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
ENV LISTEN_ADDR=":1991"
CMD ["air", "-c", ".air.toml"]
