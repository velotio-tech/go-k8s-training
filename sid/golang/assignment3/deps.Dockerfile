FROM golang:1.15 AS dep
# Add the module files and download dependencies.
ENV GO111MODULE=on
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum
WORKDIR /go/src/app
RUN go mod download
# Add the shared packages.
