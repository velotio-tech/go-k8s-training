### Run Nginx
- docker run -d -p 8080:80 nginx
- ![Nginx](./.extrafiles/nginx.png)

### Add Velotio's Page
- docker run -ti -p 8080:80 nginx bash
- - apt-get update
- - apt-get install vi
- - vi /usr/share/nginx/html/index.html
- - change the index.html file
- - exit
- docker commit `container-id` nginx:velotio
- ![Velotio](./.extrafiles/velotio.png)

### Run MySQL Container

- docker run -d --name db -e MYSQL_ROOT_PASSWORD= mysql

- docker exec -it db mysql -u root -p

- ![datbase](./.extrafiles/sql1.png)

- ![database](./.extrafiles/sql2.png)
￼
- docker run --name db -v /var/lib/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=secret -d mysql

- ![database](./.extrafiles/sql4.png)

- ![database](./.extrafiles/sql5.png)

- ![database](./.extrafiles/sql6.png)

- ![database](./.extrafiles/sql7.png)

### Containerised Go app `hello:velotio`

- ![go-app](./.extrafiles/container-app.png)

### Push the image to dockerhub

- docker tag hello:velotio jshiwam/hello:velotio
- docker push jshiwam/hello:velotio
```
The push refers to repository [docker.io/jshiwam/hello]
6b17e7c03652: Pushed 
9eed91148abf: Pushed 
4bfaf74c209f: Pushed 
fda5c4cebde5: Pushed 
e6ee11abf060: Pushed 
34a8ca206bdf: Pushed 
5543070dee1f: Pushed 
ded7a220bb05: Pushed 
velotio: digest: sha256:234c6a64275ff2f708e4621b4ee45f43b5ff079337c542296f6a27ff29f47ee3 size: 1989
```
- docker pull jshiwam/hello:velotio
