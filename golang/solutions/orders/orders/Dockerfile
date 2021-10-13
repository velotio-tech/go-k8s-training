FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .

RUN go build && \
    chmod 777 orders

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/orders /usr/bin/orders-webapp

EXPOSE 8081

ENTRYPOINT ["/usr/bin/orders-webapp"]
