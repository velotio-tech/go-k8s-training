package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

var newUser User

func Db_connectivity() *sql.DB {
	// add data into database
	db, err := sql.Open("mysql", "root:root123@tcp(127.17.0.1:3306)/userdb")
	if err != nil {
		fmt.Println("Error during connection establishment!", err)
	}
	return db
}

func CreateUser(context *gin.Context) {

	err := context.BindJSON(&newUser)
	if err != nil {
		fmt.Println("Error occured during CreateUser", err)
		return
	}

	fmt.Println("User details:")
	fmt.Println("username:", newUser.Username, "\nname:", newUser.Name, "\nemail:", newUser.Email)

	db := Db_connectivity()

	// wrapping here into string
	_, err = db.Exec("insert into user VALUES ('" + newUser.Username + "', '" + newUser.Name + "', '" + newUser.Email + "')")

	if err != nil {
		fmt.Println("Sorry, unable to insert data into table, Please try again", err)
	}

	defer db.Close()

	fmt.Println("Data added successfully!")
	context.IndentedJSON(http.StatusCreated, newUser)
}

func GetUser(context *gin.Context) {
	// get specific record from database
	uname := context.Param("username")
	str := strings.Split(uname, "=")
	uname = str[1]
	fmt.Println(uname)

	db := Db_connectivity()

	err := db.QueryRow("select * from user where username = ?", uname).Scan(&newUser.Username, &newUser.Name, &newUser.Email)
	if err != nil {
		fmt.Println("Getting error from db", err)
		context.IndentedJSON(http.StatusNotFound, newUser)
	} else {
		fmt.Println("Got the row")
		context.IndentedJSON(http.StatusOK, newUser)
	}
}

func GetAllUsers(context *gin.Context) {
	db := Db_connectivity()

	rows, _ := db.Query("select * from user")

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		err := rows.Scan(&newUser.Username, &newUser.Name, &newUser.Email)
		if err != nil {
			fmt.Println("Error while getting rows, ", err)
		}
		fmt.Println("Data received : ", newUser)
		context.IndentedJSON(http.StatusOK, newUser)
	}

	defer db.Close()
}

func DeleteUser(context *gin.Context) {
	// get the username
	uname := context.Param("username")
	str := strings.Split(uname, "=")
	uname = str[1]
	fmt.Println(uname)

	db := Db_connectivity()

	var userData User
	res, err := db.Exec("DELETE from user where username = ?", uname)

	if err != nil {
		fmt.Println("Error while getting rows, ", err)
		context.IndentedJSON(http.StatusNotFound, userData)
	} else {
		no, _ := res.RowsAffected()
		fmt.Println("Rows affected :", no)

		context.IndentedJSON(http.StatusOK, "Request success!")
	}

	defer db.Close()
}

func GetUserMeta(context *gin.Context) {
	// get the username
	uname := context.Param("username")
	str := strings.Split(uname, "=")
	uname = str[1]
	fmt.Println(uname)

	db := Db_connectivity()

	var userData User
	err := db.QueryRow("select * from user where username = ?", uname).Scan(&userData.Username, &userData.Name, &userData.Email)
	if err != nil {
		fmt.Println("Error while getting rows, ", err)
		context.IndentedJSON(http.StatusNotFound, userData)
	} else {
		fmt.Println("getting:", userData.Name)
		context.IndentedJSON(http.StatusOK, userData)
	}

	defer db.Close()
}
