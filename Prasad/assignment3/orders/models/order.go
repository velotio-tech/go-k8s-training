package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserId uint `gorm:"not null" json:"userId"`
}

func (order *Order) Create() error {
	err := GetDB().Create(order).Error
	return err
}

func (order *Order) GetUserOrders(userId uint) ([]Order, error) {
	var userOrders []Order
	err := GetDB().Where("user_id=?", userId).Find(&userOrders).Error
	return userOrders, err
}

func (order *Order) Delete(orderId uint) error {
	delOrder := &Order{}
	err := GetDB().Where("id=?", orderId).Find(delOrder).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return GetDB().Delete(delOrder).Error
}

func (order *Order) DeleteUserOrders(userId uint) error {
	var orders []Order
	return GetDB().Where("user_id=?", userId).Find(&orders).Delete(&orders).Error
}
