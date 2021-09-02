FROM golang:1.16-alpine

ENV GO111MODULE off

## Directory within image
RUN mkdir /app

## copy everything in the root directory
## into our /app directory
COPY ./app/*.go /app

WORKDIR /app

## compile the program
RUN go build

## Execute the binary
ENTRYPOINT ["/app/app"]
CMD ["--input", "Hello World!!"]
