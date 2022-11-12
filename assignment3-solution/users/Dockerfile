FROM golang:alpine3.16

COPY . /userbackend

WORKDIR /userbackend

RUN go build -o main && \
    chmod -R 777 /userbackend

CMD ["./main"]