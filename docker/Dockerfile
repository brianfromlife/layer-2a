# syntax=docker/dockerfile:1

FROM golang:1.18-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /auth ./cmd/l2a

EXPOSE 50055

VOLUME [ "/sqlite3" ]

CMD [ "/auth" ]