FROM golang:1.15.7-buster
ENV APP_USER app
ENV APP_HOME /go/src/service_order
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o service_order
ENV APP_DB_HOST mypostgres
ENV APP_DB_PORT 5432
ENV APP_DB_USERNAME postgres
ENV APP_DB_PASSWORD postgres
ENV APP_DB_NAME postgres
EXPOSE 8011
CMD ["./service_order"]
