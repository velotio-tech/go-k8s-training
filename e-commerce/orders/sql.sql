create database ecommerce;
CREATE TABLE users (
    -> userID int,
    -> user_name varchar(20),
    -> PRIMARY KEY (userID)
    -> );


insert into users values ('1', 'Jesse Pinkman');
insert into users values (2, 'Walter White');

select * from users;
+--------+---------------+
| userID | user_name     |
+--------+---------------+
|      1 | Jesse Pinkman |
|      2 | Walter White  |
+--------+---------------+
2 rows in set (0.00 sec)


create table product ( productID int, Name varchar(20), Description varchar(250), Price int, PRIMARY KEY (productID) );

 insert into product values ('1', 'Redmi4', '3GB RAM and 32 GB Storage','8999');

create table orders ( order_id int NOT NULL, product_id int NOT NULL, user_id int NOT NULL, PRIMARY KEY (order_id) );

