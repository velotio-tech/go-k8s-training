FROM golang:1.16-alpine

WORKDIR /app

COPY *.go .

RUN go build main.go

ENTRYPOINT ["./main"]