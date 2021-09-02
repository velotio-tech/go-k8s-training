# Commands for the assignment

```docker
NGINX assignment:
1. docker run --name=docker-nginx nginx.
2. docker run --name=docker-nginx -p 8080:80 -v ~/docker:/usr/share/nginx/html nginx

MYSQL assignment:
1. docker create -v /var/lib/mysql --name db-vol mysql
2. docker run -d --name mysql-db-2 --volumes-from=db-vol -e MYSQL_ROOT_PASSWORD=password -p 3307:3306 mysql
3. docker exec -it mysql-db-2 mysql -uroot -p
4. docker run -d --name mysql-db-3 --volumes-from=db-vol -e MYSQL_ROOT_PASSWORD=password -p 3308:3306 mysql
5. docker exec -it mysql-db-3 mysql -uroot -p
```

# Docker Image
- Docker hub image link - https://hub.docker.com/repository/docker/psdudeman39/hello-velotio_prasad
