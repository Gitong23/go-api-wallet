FROM golang:1.21-alpine as build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/go-sample .

# ====================

FROM alpine:3.16.2
COPY --from=build-base /app/out/go-sample /app/go-sample

# Set environment variables
ENV POSTGRES_DB_HOST=127.0.0.1 \
    POSTGRES_DB_PORT=5432 \
    POSTGRES_DB_USER=root \
    POSTGRES_DB_PASSWORD=password \
    POSTGRES_DB_NAME=wallet

CMD ["/app/go-sample"]