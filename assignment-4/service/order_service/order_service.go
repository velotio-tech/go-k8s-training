package orderservice

import (
	"github.com/practice/model"
	"github.com/practice/repository"
)

func CreateNewOrder(username string, order model.Order) (model.Order, error) {
	order.UserName = username
	createdOrder, err := repository.CreateNewOrderByUserName(order)
	return createdOrder, err
}

func GetOrderByUserNameOrderID(username string, orderID int) (model.Order, error) {
	order, err := repository.GetOrderByUserNameAndOrderID(username, orderID)
	return order, err
}

func GetAllOrders(username string) ([]model.Order, error) {
	orders, err := repository.GetAllOrdersByUserName(username)
	return orders, err
}

func UpdateOrder(order model.Order) (model.Order, error) {

	_, err := repository.GetOrderByUserNameAndOrderID(order.UserName, order.OrderID)
	if err != nil {
		return model.Order{}, err
	}
	updatedOrder, err := repository.UpdateUserByUserNameAndOrderID(order)
	return updatedOrder, err
}

func DeleteOrder(username string, orderID int) error {
	_, err := repository.GetOrderByUserNameAndOrderID(username, orderID)
	if err != nil {
		return err
	}
	err = repository.DeleteOrderByUserNameAndOrderID(username, orderID)
	return err
}
