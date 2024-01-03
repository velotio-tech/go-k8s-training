package repository

import (
	"time"

	"github.com/pratikpjain/go-k8s-training/assignment3/db"
	model "github.com/pratikpjain/go-k8s-training/assignment3/models"
)

const (
	getOrderByUserNameAndOrderID    = `SELECT order_id, order_desc, username, amount, created_at, updated_at FROM ORDERS WHERE username=$1 AND order_id=$2`
	getAllOrdersByUserName          = `SELECT order_id, order_desc, username, amount, created_at, updated_at FROM ORDERS WHERE username=$1`
	insertOrderByUserName           = `INSERT INTO ORDERS (order_id, order_desc, username, amount, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING order_id, order_desc, username, amount, created_at, updated_at`
	updateOrderByUserNameAndOrderID = `UPDATE ORDERS SET order_desc = $3, amount = $4, updated_at = now() where username = $1 AND order_id = $2 RETURNING order_id, order_desc, username, amount, created_at, updated_at`
	deleteOrderByUserNameAndOrderID = `DELETE FROM ORDERS WHERE username = $1 AND order_id = $2`
)

func DeleteOrderByUserNameAndOrderID(username string, orderID int) error {
	_, err := db.GetDB().Exec(deleteOrderByUserNameAndOrderID, username, orderID)
	return err
}

func UpdateUserByUserNameAndOrderID(order model.Order) (model.Order, error) {

	var updatedOrder model.Order

	err := db.GetDB().QueryRow(updateOrderByUserNameAndOrderID, order.UserName, order.OrderID, order.OrderDesc, order.Amount).Scan(&updatedOrder.OrderID, &updatedOrder.OrderDesc, &updatedOrder.UserName, &updatedOrder.Amount, &updatedOrder.CreatedAt, &updatedOrder.UpdatedAt)
	if err != nil {
		return model.Order{}, err
	}

	return updatedOrder, nil
}

func CreateNewOrderByUserName(order model.Order) (model.Order, error) {

	var orderCreated model.Order

	err := db.GetDB().QueryRow(insertOrderByUserName, order.OrderID, order.OrderDesc, order.UserName, order.Amount).Scan(&orderCreated.OrderID, &orderCreated.OrderDesc, &orderCreated.UserName, &orderCreated.Amount, &orderCreated.CreatedAt, &orderCreated.UpdatedAt)
	if err != nil {
		return model.Order{}, err
	}

	return orderCreated, nil
}

func GetAllOrdersByUserName(username string) ([]model.Order, error) {

	var orders []model.Order
	var orderDesc string
	var amount float64
	var orderID int
	var created_at, updated_at time.Time

	rows, err := db.GetDB().Query(getAllOrdersByUserName, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&orderID, &orderDesc, &username, &amount, &created_at, &updated_at)
		if err != nil {
			return nil, err
		}
		order := model.Order{
			OrderID:   orderID,
			OrderDesc: orderDesc,
			UserName:  username,
			Amount:    amount,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByUserNameAndOrderID(username string, orderID int) (model.Order, error) {

	var orderDesc string
	var amount float64
	var created_at, updated_at time.Time

	row := db.GetDB().QueryRow(getOrderByUserNameAndOrderID, username, orderID)
	err := row.Scan(&orderID, &orderDesc, &username, &amount, &created_at, &updated_at)
	if err != nil {
		return model.Order{}, err
	}
	order := model.Order{
		OrderID:   orderID,
		OrderDesc: orderDesc,
		UserName:  username,
		Amount:    amount,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
	return order, nil

}
