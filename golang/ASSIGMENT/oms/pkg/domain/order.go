package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"oms/pkg/models"
)

// CreateOrder ...
func (o *OrderCliet) CreateOrder(order *models.Order) (string, error) {
	result, err := o.DB.Exec(`INSERT INTO orders ("ID", "buyerUserID", status, "totalAmount", products, "paymentID")
	VALUES ($1, $2, $3, $4, $5, $6)`, order.ID, order.BuyerUserID, order.Status, order.TotalAmount, order.Products, order.PaymentID)
	if err != nil {
		log.Println("failed to create order : ", err)
		return "", ErrOrderCreationFailed
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		log.Println("failed to retirve rows affected data : ", err)
		return "", ErrOrderCreationFailed
	}

	return fmt.Sprint(result.LastInsertId()), nil
}

// GetOrder gets a single order using orderID
func (o *OrderCliet) GetOrder(orderID string) (*models.Order, error) {
	var err error
	order := &models.Order{}

	getOrderQuery := `SELECT ""ID","buyerUserID", status, "totalAmount", products, "paymentID", "createdAt","updatedAt" FROM "orders" WHERE "ID"=$1'`
	err = o.DB.QueryRow(getOrderQuery, orderID).Scan(
		&order.ID,
		&order.BuyerUserID,
		&order.Status,
		&order.TotalAmount,
		&order.Products,
		&order.PaymentID,
		&order.CreatedAt,
		&order.UpdatedAt)

	if err != nil {
		log.Println("failed to retrive the user : ", err)
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrOrderNotFound
		}
		return nil, ErrInternalError
	}

	return order, nil
}

// GetOrders gets all the orders of a particular buyerUserID
func (o *OrderCliet) GetOrdersByUserID(userID string) ([]models.Order, error) {
	orders := []models.Order{}

	getOrderQuery := `SELECT ""ID","buyerUserID", status, "totalAmount", products, "paymentID", "createdAt","updatedAt" FROM "orders" WHERE "buyerUserID"=$1'`
	rows, err := o.DB.Query(getOrderQuery, userID)
	if err != nil {
		log.Println("failed to fetch orders : ", err)
		if errors.Is(err, sql.ErrNoRows) {
			return orders, nil
		}
		return nil, ErrInternalError
	}
	for rows.Next() {
		order := models.Order{}
		rows.Scan(
			&order.ID,
			&order.BuyerUserID,
			&order.Status,
			&order.TotalAmount,
			&order.Products,
			&order.PaymentID,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		orders = append(orders, order)
	}

	return orders, nil
}

// DeleteOrderByUserID deletes all orders of a user from the table
func (o *OrderCliet) DeleteOrderByUserID(userID string) error {
	deleteOrdersQuery := `DELETE * from "orders" where "buyerUserID"=$1;`
	result, err := o.DB.Exec(deleteOrdersQuery, userID)
	if err != nil {
		log.Println("failed to delete all the orders : ", err)
		return ErrInternalError
	}
	if rowsEffected, err := result.RowsAffected(); err != nil || rowsEffected == 0 {
		return ErrFailedToDeleteOrders
	}

	return nil
}

// DeleteOrderByOrderID deletes a specific order of a user from the table
func (o *OrderCliet) DeleteOrderByOrderID(orderID string) error {
	deleteOrdersQuery := `DELETE * from "orders" where "ID"=$1;`
	result, err := o.DB.Exec(deleteOrdersQuery, orderID)
	if err != nil {
		log.Println("failed to delete  the specified order : ", err)
		return ErrFailedToDeleteOrder
	}
	if rowsEffected, err := result.RowsAffected(); err != nil || rowsEffected == 0 {
		return ErrFailedToDeleteOrder
	}

	return nil
}
