FROM golang:alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o users .

FROM alpine
COPY --from=builder /app/users /app/
WORKDIR /app
ENTRYPOINT ["./users"]

