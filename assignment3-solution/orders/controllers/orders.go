package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	OrderId   int64  `json:"order_id"`
	OrderName string `json:"order_name"`
	Price     int64  `json:"price"`
}

var newOrder Order

func Db_connectivity() *sql.DB {
	// add data to database
	db, err := sql.Open("mysql", "root:root123@tcp(127.17.0.1:3306)/userdb")
	if err != nil {
		fmt.Println("Error during db connection", err)
	}

	return db
}

func CreateOrders(context *gin.Context) {

	uname := context.Param("username")

	err := context.BindJSON(&newOrder)
	if err != nil {
		fmt.Println("Error while loading data!", err)
	}

	// send this to calling function and it will store it in the database
	fmt.Println("uname: ", uname, newOrder.OrderId, newOrder.OrderName, newOrder.Price)

	db := Db_connectivity()

	// wrapping here into string
	query := "insert into orders VALUES (" + strconv.Itoa(int(newOrder.OrderId)) + ", '" + newOrder.OrderName + "', " + strconv.Itoa(int(newOrder.Price)) + ", '" + uname + "')"
	_, err = db.Exec(query)

	if err != nil {
		fmt.Println("Sorry, unable to insert data into table, Please try again", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		fmt.Println("Data added successfully!")
		context.IndentedJSON(http.StatusCreated, newOrder)
	}

	defer db.Close()
}

func GetAllOrders(context *gin.Context) {
	uname := context.Param("username")

	db := Db_connectivity()

	rows, _ := db.Query("select * from orders where username = ?", uname)

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		err := rows.Scan(&newOrder.OrderId, &newOrder.OrderName, &newOrder.Price, &uname)
		if err != nil {
			fmt.Println("Error while fetching records!", err)
		}
		context.JSON(http.StatusOK, newOrder)
	}
	defer db.Close()
}

func GetOrderByOrderId(context *gin.Context) {
	uname := context.Param("username")
	oid := context.Param("order_id")

	db := Db_connectivity()

	err := db.QueryRow("select * from orders where order_id = ? and username=?", oid, uname).Scan(&newOrder.OrderId, &newOrder.OrderName, &newOrder.Price, &uname)
	if err != nil {
		fmt.Println("Getting error from db", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		fmt.Println("Got the row")
		context.IndentedJSON(http.StatusOK, newOrder)
	}
}

func UpdateOrder(context *gin.Context) {
	uname := context.Param("username")
	oid := context.Param("order_id")
	err := context.BindJSON(&newOrder)
	if err != nil {
		fmt.Println("Getting error while loading json from user")
		return
	}

	fmt.Println(uname, oid, newOrder.OrderName)

	db := Db_connectivity()

	rows, err := db.Query("update orders set order_name=? where order_id=? AND username=?", newOrder.OrderName, oid, uname) //.Scan(&newOrder.OrderId, &newOrder.OrderName, &newOrder.Price, &uname)
	if err != nil {
		fmt.Println("Getting error from db", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		fmt.Println("Updated the row", rows)
		context.IndentedJSON(http.StatusOK, "Row updated!")
	}
}

func DeleteAllOrders(context *gin.Context) {

	uname := context.Param("username")

	db := Db_connectivity()

	rows, err := db.Query("delete from orders where username=?", uname)
	// Loop through rows, using Scan to assign column data to struct fields.

	if err != nil {
		fmt.Println("Error while fetching records!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		fmt.Println("Total rows affected : ", rows)
		context.IndentedJSON(http.StatusOK, "Record deleted successfully!")
	}
	defer db.Close()
}

func DeleteOrder(context *gin.Context) {
	uname := context.Param("username")
	oid := context.Param("order_id")

	db := Db_connectivity()

	rows, err := db.Query("delete from orders WHERE order_id = ? AND username=?", oid, uname)

	if err != nil {
		fmt.Println("Error while fetching records!", err)
		context.IndentedJSON(http.StatusNotFound, err)
	} else {
		fmt.Println("Total rows affected : ", rows)
		context.IndentedJSON(http.StatusOK, "Record deleted successfully!")
	}
	defer db.Close()
}
