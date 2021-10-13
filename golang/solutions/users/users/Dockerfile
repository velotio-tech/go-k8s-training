FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .

RUN go build && \
    chmod 777 users

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/users /usr/bin/users-webapp

EXPOSE 8080

ENTRYPOINT ["/usr/bin/users-webapp"]
