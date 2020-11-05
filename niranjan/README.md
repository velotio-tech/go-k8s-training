
# E-Commerce Application

## Introduction

This is a simple E-commerce application created as an exercise to improve my existing knowledge of GoLang, Docker, Kubernetes & Helm.

This application consists of three microservices: Users-ms, Orders-ms, database-ms

### Users-ms:
Responsible for Creating, Reading, Updating & Deleting user data on the database.

### Orders-ms:
Responsible for Creating, Reading, Updating & Deleting orders of specific user on the database.

### Database-ms:
MySQL database is used for this application. Other microservices interact with this db using GORM.

------

## Getting Started

### Prerequisites:

* minikube
* kubectl
* helm
* go
* docker

### Environment Setup:

Create a local minikube cluster
`minikube start`

Use local registry with minikube
`eval $(minikube docker-env)`

Pull MySQL image which we will be using for our database
`docker pull mysql`

### Dockerize the microservices

Dockerize users microservice
`cd users`
`docker build -t users:1.0 .`

Dockerize orders microservice
`cd ../orders`
`docker build -t orders:1.0 .`

### Deploying the application on minikube using helm

`cd ..`

Check if the templates created by helm have all the right values
`helm template myapp/`

To identify any issues before installing the helm chart
`helm lint myapp/`

To check is anything is wrong with helm chart configuration
`helm install myapp-release --debug --dry-run myapp/`

Install the helm chart
`helm install myapp-release myapp/`

List & check helm release details
`helm list -a`

Application should be running in minikube now!

### Accessing the application

Get the IP address to minikube master node
`kubectl cluster-info`

Use node-port 30111 with the masterIP to make requests to our application

### Request Example

### Requests handled by USER microservice

GET     http://masterIP:30111/users     -->     GET all users

GET     http://masterIP:30111/users/UserID     -->     GET specific user data

POST     http://masterIP:30111/users     -->     POST user data
```
{
    "name": "niranjan",
    "email": "niranjan@xyz.xom",
    "password": "niranjan"
}
```

PUT     http://masterIP:30111/users/UserID     -->     Update user data
```
{
    "name": "Shyam",
    "email": "shyam@xyz.xom",
    "password": "shyam"
}
```

DELETE     http://masterIP:30111/users/UserID     -->     DELETE user


### Requests handled by USER microservice

GET     http://masterIP:30111/users/UserID     -->     GET specific users all orders

GET     http://masterIP:30111/users/UserID/orders/OrderID     -->     GET specific users, specific order

POST     http://masterIP:30111/users/UserID     -->     POST user order
```
{
    "bill_amount": 5000
}
```

PUT     http://masterIP:30111/users/UserID/orders/OrderID     -->     Update order data
```
{
    "bill_amount": 7000
}
```

DELETE     http://masterIP:30111/users/UserID/orders    -->     DELETE specific users all orders

DELETE     http://masterIP:30111/users/UserID/orders/OrderID     -->     DELETE specific users, specific order