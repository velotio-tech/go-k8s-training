FROM golang:1.16

WORKDIR /app

COPY . .

RUN go build -o userapi api.go

EXPOSE 8000

CMD ./userapi
