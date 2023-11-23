# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from the latest golang base image
FROM golang:1.21.4-alpine3.17 AS builder


RUN apk update
RUN apk add git gcc musl-dev
RUN apk add curl bash

WORKDIR /go/src/app

# Copy the Go application source code into the container
COPY . .

# Add Maintainer Info
LABEL maintainer="Arijit Nayak"
ENV APP_NAME "oms"

#You can define a variable with:
#ARG myvalue=3
# Create work env
RUN mkdir -p  $GOPATH/src/$CODE_PATH/
ADD .  $GOPATH/src/$CODE_PATH/
WORKDIR $GOPATH/src/$CODE_PATH/

# Run go clean?
RUN go install -tags dynamic ./...

# run the command to start es data service
CMD $GOPATH/bin/ums

# Expose port 8082 to the outside world
EXPOSE 5060