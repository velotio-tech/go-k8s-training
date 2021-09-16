package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	host := os.Getenv("E_COMM_DB_HOST")
	user := os.Getenv("E_COMM_DB_USER")
	dbName := os.Getenv("E_COMM_DB_NAME")
	passwd := os.Getenv("E_COMM_DB_PASS")
	port := os.Getenv("E_COMM_DB_PORT")
	dbConnStr := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, passwd, port, dbName)

	conn, err := gorm.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{})

	log.Println("Connected to database successfully...")
}

func GetDB() *gorm.DB {
	return db
}
