FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

EXPOSE ${PORT}

CMD ["./main"]