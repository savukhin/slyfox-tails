FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /main .

EXPOSE $BACKEND_PORT

CMD ["/main"]
