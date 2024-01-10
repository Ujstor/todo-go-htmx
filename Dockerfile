FROM golang:1.21-alpine as base

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base as dev
RUN go get -u github.com/cosmtrek/air
RUN make build
EXPOSE ${PORT}
CMD [ "make", "watch" ]


FROM base as migrate
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
CMD ["goose", "-dir", "./migrations", "up"]