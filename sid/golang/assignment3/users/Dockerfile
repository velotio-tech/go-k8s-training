FROM dependencies AS builder
# Copy the application source code.
COPY . /go/src/app/users
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/users /go/src/app/users/
ENTRYPOINT [ "/go/bin/users" ]

FROM ubuntu:latest
COPY --from=builder /go/bin/users /bin/users
ENTRYPOINT [ "/bin/users" ]
