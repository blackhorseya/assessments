FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o app main.go

FROM alpine:3 AS final

WORKDIR /app

COPY --from=builder /src/app ./

ENTRYPOINT ./app
