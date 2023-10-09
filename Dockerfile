FROM golang:latest

WORKDIR /app

ENV CACHE_SIZE=10

COPY . .

RUN go build -o main.go .

EXPOSE 8080

CMD ["./main.go"]
