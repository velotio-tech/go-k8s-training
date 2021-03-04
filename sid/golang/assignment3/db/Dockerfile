FROM dependencies AS builder

COPY . /go/src/app/db

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/db /go/src/app/db/
ENTRYPOINT [ "/go/bin/db" ]

FROM ubuntu:latest
COPY --from=builder /go/bin/db /bin/db
ENTRYPOINT [ "/bin/db" ]
