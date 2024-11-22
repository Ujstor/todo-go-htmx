FROM golang:1.22.1-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.19.0 AS prod
WORKDIR /app
COPY --from=base /app/main /app/main
COPY cmd/web cmd/web
EXPOSE ${PORT}
CMD ["./main"]
