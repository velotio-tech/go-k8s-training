# assignment 3 backend service for e commerse website

# Overview

The backend consists of 3 services
1. service_order - crud functionality of order
2. service_user - crud functionality of user as well as calling service order internally
3. postgres db service - holds users and orders data

service_user is only accessible from outside and all the operations on users and orders will be done via this service only

# Installation
1. install docker

2. use docker-compose for running with docker
```cmd 
$ docker-compose up --build
```

# API Reference

Api will be available at http://localhost:8010
Consumes json produces json


# 1. Get list of users

**URL** : `/users`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Response example**

```json
[
{
        "ID": 1,
        "CreatedAt": "2021-03-04T20:15:17.27306Z",
        "UpdatedAt": "2021-03-04T20:15:17.27306Z",
        "DeletedAt": null,
        "Name": "swapnil s",
        "Email": "swapnils@example.com"
    }
]
```

# 2. Get a user

**URL** : `/users/<id:int>`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`


**Response example**

```json
{
        "ID": 1,
        "CreatedAt": "2021-03-04T20:15:17.27306Z",
        "UpdatedAt": "2021-03-04T20:15:17.27306Z",
        "DeletedAt": null,
        "Name": "swapnil s",
        "Email": "swapnils@example.com"
    }
```


# 3. Create a user

**URL** : `/users`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `201 Created`

**Request example**

```json
{        
    "Name": "swapnil s",
    "Email": "swapnils@example.com"
}
```


**Response example**

```json
{
        "ID": 1,
        "CreatedAt": "2021-03-04T20:15:17.27306Z",
        "UpdatedAt": "2021-03-04T20:15:17.27306Z",
        "DeletedAt": null,
        "Name": "swapnil s",
        "Email": "swapnils@example.com"
    }
```

## Notes

* name and email is required


# 4. Update a user

**URL** : `/users/<id:int>`

**Method** : `PUT`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Request example**

```json
{        
    "Name": "swapnil s",
    "Email": "swapnils@example.com"
}
```


**Response example**

```json
{
        "ID": 1,
        "CreatedAt": "2021-03-04T20:15:17.27306Z",
        "UpdatedAt": "2021-03-04T20:15:17.27306Z",
        "DeletedAt": null,
        "Name": "swapnil s",
        "Email": "swapnils@example.com"
    }
```

## Notes

* name and email is required

# 5. Delete a user

**URL** : `/users/<id:int>`

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`


**Response example**

```json
{
    "result": "success"
}
```

# 6. Get list of orders of a user

**URL** : `/users/<userid:int>/orders`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Response example**

```json
{
    "orders": [
        {
            "CreatedAt": "2021-03-04T20:16:01.558815Z",
            "DeletedAt": null,
            "ID": 4,
            "Name": "vadapav",
            "Quantity": 1,
            "Unit": "qty",
            "UpdatedAt": "2021-03-04T20:16:01.558815Z",
            "UserID": 6
        },
        {
            "CreatedAt": "2021-03-04T20:16:01.558815Z",
            "DeletedAt": null,
            "ID": 5,
            "Name": "samosa",
            "Quantity": 1,
            "Unit": "qty",
            "UpdatedAt": "2021-03-04T20:16:01.558815Z",
            "UserID": 6
        }
    ]
}
```
# 7. Get a order

**URL** : `/users/<userid:int>/orders/<orderid:int>`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`


**Response example**

```json
{
    "CreatedAt": "2021-03-04T20:16:01.558815Z",
    "DeletedAt": null,
    "ID": 6,
    "Name": "tea",
    "Quantity": 1,
    "Unit": "cups",
    "UpdatedAt": "2021-03-04T20:16:01.558815Z",
    "UserID": 6
}
```

# 8. Create a order

**URL** : `/users/<userid:int>/orders`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `201 Created`

**Request example**

```json
{        
    "Name": "samosa",
    "quantity": 5,
    "unit": "pieces"
}
```


**Response example**

```json
{
    "CreatedAt": "2021-03-04T20:16:01.558815Z",
    "DeletedAt": null,
    "ID": 6,
    "Name": "samosa",
    "Quantity": 5,
    "Unit": "pieces",
    "UpdatedAt": "2021-03-04T20:16:01.558815Z",
    "UserID": 6
}
```

## Notes

* in current version one order can only have one item.

# 9. Update a order

**URL** : `/users/<userid:int>/orders/<orderid:int>`

**Method** : `PUT`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Request example**

```json
{        
    "Name": "samosa",
    "quantity": 6,
    "unit": "pieces"
}
```


**Response example**

```json
{
    "CreatedAt": "2021-03-04T20:16:01.558815Z",
    "DeletedAt": null,
    "ID": 6,
    "Name": "samosa",
    "Quantity": 6,
    "Unit": "pieces",
    "UpdatedAt": "2021-03-04T20:16:01.558815Z",
    "UserID": 6
}
```

# 10. Delete a order

**URL** : `/users/<userid:int>/orders/<orderid:int>`

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`


**Response example**

```json
{
    "result": "success"
}
```


# 11. Delete all orders of a user

**URL** : `/users/<userid:int>/orders`

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 OK`


**Response example**

```json
{
    "result": "success"
}
```