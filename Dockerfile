FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o coding-challenge .

CMD ["./coding-challenge"]