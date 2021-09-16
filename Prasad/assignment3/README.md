# How To Use

Run the project using docker-compose. -> ``` docker-compose up```

# API details

The user microservice application is exposed on host port - ```7000```

## User APIs

#### Create User
1.
    url: http://localhost:7000/user
    method: POST
    input-body: ``` { "email": "prasad.saraf@velotio.com", "name": "Prasad Saraf" } ```
    
#### Delete User
2.
    url: http://localhost:7000/user
    method: DELETE
    input-body: ``` { "id": 13 } ```

#### Get all users
3.
    url: http://localhost:7000/users
    method: GET
    input-body: none

#### Update user
4.
    url: http://localhost:7000/user
    method: PUT
    input-body: ``` { "ID": 1, "email": "prasad.saraf@velotio.com", "name": "Prasad Saraf" } ```

#### User exists
5.
    url: http://localhost:7000/user/exists/```{userId}```
    method: GET
    input-body: none

## User-Order APIs

#### Create user order
1.
    url: http://localhost:7000/user/order
    method: POST
    input-body: ```{ "userId": 1 }```

#### Get all the orders of user
2.
    url: http://localhost:7000/user/```{userId}```/orders
    method: GET
    input-body: none

#### Delete all orders of user
3.
    url: http://localhost:7000/user/```{userId}```/orders
    method: DELETE
    input-body: none

#### Delete order of user
4.
    url: http://localhost:8090/user/```{userId}```/order/```{orderId}```
    method: DELETE
    input-body: none