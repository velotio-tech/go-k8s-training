FROM golang:1.16

WORKDIR /app

COPY . .

RUN go build -o orderapi api.go

EXPOSE 8001

CMD ./orderapi
