FROM golang:1.17-alpine

ENV GO111MODULE auto

ENV APP_HOME go/src/orders
WORKDIR $APP_HOME

COPY . .

ENV E_COMM_DB_HOST postgres_cont
ENV E_COMM_DB_PASS password
ENV E_COMM_DB_USER postgres
ENV E_COMM_DB_PORT 5432
ENV E_COMM_DB_NAME postgres
ENV USER_SERVER_PORT 8090
ENV USER_SERVER_BASE_URL http://user-app:8090/
ENV USER_SERVER_EXISTS_URL user/exists/

RUN go mod verify
RUN go mod download
RUN go mod tidy
RUN go build
CMD ["./orders"]