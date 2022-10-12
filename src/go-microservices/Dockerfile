FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
# EXPOSE 8080
# ENTRYPOINT ["./main"]


FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/main"]