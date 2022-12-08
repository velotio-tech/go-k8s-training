package controller

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Connect() *gorm.DB {
	db, err = gorm.Open("mysql", "root:Test@123@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("Connection Failed to Open")
		panic(err)
	}
	return db
}
