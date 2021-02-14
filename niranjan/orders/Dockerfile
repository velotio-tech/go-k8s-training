FROM golang:alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o orders .

FROM alpine
COPY --from=builder /app/orders /app/
WORKDIR /app
ENTRYPOINT ["./orders"]
