FROM golang:1.14.7-alpine3.12

WORKDIR /app

COPY . .

RUN go build -o build/main .

CMD ["./build/main"]
