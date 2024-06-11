FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main ./app

EXPOSE 8080

CMD ["./main"]
