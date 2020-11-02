package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	// blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Order is a struct containing Order table fields
type Order struct {
	UserID     uint64    `gorm:"ForeignKey:ID" json:"user_id"`
	OrderID    uint64    `gorm:"primary_key;auto_increment" json:"order_id"`
	BillAmount uint64    `gorm:"not null" json:"bill_amount"`
	CreatedAt  time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// TableCreate Creates new table schema for User table
func TableCreate(l *log.Logger) {
	db := Connect()
	defer db.Close()
	//log.Println("dropping table!")
	//db.Debug().DropTableIfExists(&Order{})
	log.Println("Creating/Updating table!")
	db.Debug().AutoMigrate(&Order{})
}

// db credential information
const (
	USER   = "niranjan"
	PASS   = "niranjan"
	HOST   = "example"
	PORT   = 3306
	DBNAME = "velotio"
)

// Connect uses the above mentioned creds to connect to the velotio DB
func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

// CreateUserOrder creates order for specific user  - C
func CreateUserOrder(order Order, userID uint64) error {
	db := Connect()
	defer db.Close()
	var err error
	order.UserID = userID
	if err != nil {
		return err
	}
	err = db.Create(&order).Error
	return err
}

// GetAll returns all orders of specific user in DB - R
func GetAll(userID uint64) interface{} {
	db := Connect()
	defer db.Close()
	return db.Where("user_id = ?", userID).Find(&[]Order{}).Value
}

// GetByOrderID returns specific order of specific user in DB - R
func GetByOrderID(userID uint64, orderID uint64) interface{} {
	db := Connect()
	defer db.Close()
	return db.Where("user_id = ? AND order_id = ?", userID, orderID).Find(&Order{}).Value
}

// UpdateUserOrder Only for changing bill_amount of specific order of specific user - U
func UpdateUserOrder(order Order, userID uint64, orderID uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&order).Where("user_id = ? AND order_id = ?", userID, orderID).UpdateColumns(
		map[string]interface{}{
			"bill_amount": order.BillAmount,
		},
	)
	return rs.RowsAffected, rs.Error
}

// DeleteOrder deletes the specified order of specified user - D
func DeleteOrder(userID uint64, orderID uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	rs = db.Where("user_id = ? AND order_id = ?", userID, orderID).Delete(&Order{})
	return rs.RowsAffected, rs.Error
}

// DeleteOrders deletes all order of specified user - D
func DeleteOrders(userID uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	rs = db.Where("user_id = ?", userID).Delete(&Order{})
	return rs.RowsAffected, rs.Error
}
