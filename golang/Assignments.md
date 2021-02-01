# Assignments
---
### Assignment 1
Topics covered:
API development and Http communication using client, Json Encoding/decoding, DB connection, Docker , Helm, k8s
Create a private git repo for this and add a Readme.md explaining the steps to run you application using Docker and k8s

- Create a small e-commerce backend having two micro-services: `User` and `Order`. For now ignore the authentication part. `User` service only will be exposed to the public and will communicate with the `Order` service for any order related queries.

- You are admin and can perform all operations. Things to consider
1. Handle CRUD operation of Users and Orders. There can be 1:n mapping between user and orders.
2. Can only update one order at a time
3. Can delete one or all orders of a user at a time
4. Store the data in any DB of your choice
5. Containerize your services and test it.
6. Package your application using Helm and deploy on K8s and test it.
   
   Eg: 
   
   http://localhost:80/users -> all users
   
   http://localhost:80/users/0/orders -> all orders of user 0
   
   http://localhost:80/users/0/orders/0 - order 0 of user 0
   
---


