FROM dependencies AS builder
# Copy the application source code.
COPY . /go/src/app/orders
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/orders /go/src/app/orders/
ENTRYPOINT [ "/go/bin/orders" ]

FROM ubuntu:latest
COPY --from=builder /go/bin/orders /bin/orders
ENTRYPOINT [ "/bin/orders" ]
