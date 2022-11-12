FROM golang:alpine3.16

COPY . /orderbackend

WORKDIR /orderbackend

RUN go build -o main && \
    chmod -R 777 /orderbackend

CMD ["./main"]