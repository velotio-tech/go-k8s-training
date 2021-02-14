package models

import (
	"fmt"
	"log"
	"time"

	"example.com/users/cryptography"
	"github.com/jinzhu/gorm"

	// blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User is a struct containing User table fields
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(20);not null" json:"name"`
	Email     string    `gorm:"type:varchar(40);not null;unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(60); not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// TableCreate Creates new table schema for User table
func TableCreate(l *log.Logger) {
	log.Println(USER, PASS, HOST, PORT, DBNAME)
	db := Connect()
	defer db.Close()
	//log.Println("dropping table!")
	//db.Debug().DropTableIfExists(&User{})
	log.Println("Creating/Updating table!")
	db.Debug().AutoMigrate(&User{})
}

// db credential information
const (
	USER   = "niranjan"
	PASS   = "niranjan"
	HOST   = "example-svc"
	PORT   = 3306
	DBNAME = "velotio"
)

// Connect uses the above mentioned creds to connect to the velotio DB
func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

// CreateUser only converts passwd to hashed passwd and create user - C
func CreateUser(user User) error {
	db := Connect()
	defer db.Close()
	var err error
	user.Password, err = cryptography.Hash(user.Password)
	if err != nil {
		return err
	}
	err = db.Create(&user).Error
	return err
}

// GetAll returns all users in DB - R
func GetAll() interface{} {
	db := Connect()
	defer db.Close()
	return db.Order("id asc").Find(&[]User{}).Value
}

// GetByID returns user of specified ID - R
func GetByID(id uint64) interface{} {
	db := Connect()
	defer db.Close()
	return db.Where("id = ?", id).Find(&User{}).Value
}

// UpdateUser Only for changing name & email - U
func UpdateUser(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&user).Where("id = ?", user.ID).UpdateColumns(
		map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		},
	)
	return rs.RowsAffected, rs.Error
}

// Delete the specified user - D
func Delete(id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	rs = db.Where("id = ?", id).Delete(&User{})
	return rs.RowsAffected, rs.Error
}
