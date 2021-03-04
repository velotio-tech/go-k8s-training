FROM dependencies AS builder
# Copy the application source code.
COPY . /go/src/app/db
# Build the application.
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/db /go/src/app/db/
ENTRYPOINT [ "/go/bin/db" ]

FROM ubuntu:latest
COPY --from=builder /go/bin/db /bin/db
ENTRYPOINT [ "/bin/db" ]
