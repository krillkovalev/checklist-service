FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go build -o kafka_service

CMD ["./kafka_service"]

