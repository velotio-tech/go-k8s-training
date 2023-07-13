FROM golang:1.18-alpine
RUN mkdir /backend-app
COPY  . /backend-app
WORKDIR /backend-app
RUN go build -o main .
CMD ["/backend-app/main"]