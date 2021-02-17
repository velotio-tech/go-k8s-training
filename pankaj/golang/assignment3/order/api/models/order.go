package models

import (
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/order/api/utils"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Name   string `json:"name"`
	UserID int    `json:"userID"`
}

func (order *Order) Create() map[string]interface{} {
	err := GetDB().Create(order).Error

	if err != nil {
		return utils.Message(false, "Failed to create order, connection error.")
	}

	response := utils.Message(true, "Successfully Created Order")
	response["order"] = order
	return response
}

func (order *Order) Update(id int) map[string]interface{} {
	temp := &Order{}

	err := GetDB().Where("id=?", id).First(temp).Error
	if err != nil {

		return utils.Message(true, "Order Not Found!")
	}

	temp.Name = order.Name
	GetDB().Save(temp)

	response := utils.Message(true, "Successfully Updated Order")
	response["order"] = order
	return response
}

func GetOrder(id int) *Order {

	order := &Order{}
	err := GetDB().Where("id = ?", id).First(order).Error
	if err != nil { //Order not found!
		return nil
	}

	return order
}

func DeleteOrder(id int) map[string]interface{} {
	order := &Order{}
	err := GetDB().Where("id=?", id).Find(order).Error
	if err != nil {
		return utils.Message(false, "Order Not Found!")
	}
	GetDB().Delete(order)
	return utils.Message(true, "Successfully Deleted Order")
}

// func GetUserOrder(id int) *Order {

// }
