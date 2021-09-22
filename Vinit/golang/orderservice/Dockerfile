FROM golang:latest
WORKDIR /app
COPY . .
EXPOSE 9009
RUN go mod tidy
RUN go build -o ./orderbackend
CMD ["./orderbackend"]

