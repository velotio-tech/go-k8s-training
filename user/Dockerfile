FROM golang:alpine
WORKDIR /user
COPY . /user
RUN go build -o main
CMD [ "./main" ]