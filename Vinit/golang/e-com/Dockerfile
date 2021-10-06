FROM golang:latest
WORKDIR /app
COPY . .
EXPOSE 8009
RUN go mod tidy
RUN go build -o ./ecom
CMD ["./ecom"]
