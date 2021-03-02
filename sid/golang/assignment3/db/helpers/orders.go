package helpers

import (
	"database/sql"
	"fmt"
)

func CreateOrder(conn *sql.DB, userID int, details string) error {
	result, err := conn.Exec(fmt.Sprintf("insert into orders (details) values ('%s')", details))

	if err != nil {
		fmt.Println("Failed to create user cuz", err)
		return err
	}

	orderID, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Failed to get order ID", err)
		return err
	}

	_, err = conn.Exec(fmt.Sprintf("insert into users_orders_mapping (user_id, order_id) values (%d, %d)", userID, orderID))

	if err != nil {
		fmt.Println("Failed to create user_order mapping cuz", err)
		return err
	}

	return nil
}

func GetOrder(conn *sql.DB, id int) (map[int]string, error) {
	rows, err := conn.Query(fmt.Sprintf("select id, details from orders where id=%d", id))
	orders := make(map[int]string)

	if err != nil {
		fmt.Println("Failed to get order cuz", err)
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var details string
		var id int
		err = rows.Scan(&id, &details)

		if err != nil {
			fmt.Println("Failed to populate order cuz", err)
			return orders, err
		}

		orders[id] = details
	}

	return orders, nil
}

func UpdateOrder(conn *sql.DB, id int, details string) error {
	_, err := conn.Exec(fmt.Sprintf("update orders set details = '%s' where id = %d", details, id))

	if err != nil {
		fmt.Println("Failed to update order cuz", err)
		return err
	}

	return nil
}

func DeleteOrder(conn *sql.DB, id int) error {
	_, err := conn.Exec(fmt.Sprintf("delete from orders where id = %d", id))

	if err != nil {
		fmt.Println("Failed to delete order cuz", err)
		return err
	}

	_, err = conn.Exec(fmt.Sprintf("delete from users_orders_mapping where order_id = %d", id))

	if err != nil {
		fmt.Println("Failed to delete order cuz", err)
		return err
	}

	return nil
}

func GetOrders(conn *sql.DB, IDs []int) (map[int]string, error) {
	stmt := "select id, details from orders where id in ( "

	for _, id := range IDs {
		stmt += fmt.Sprintf("%d, ", id)
	}

	stmt = stmt[:len(stmt)-2]

	rows, err := conn.Query(stmt + " )")
	orders := make(map[int]string)

	if err != nil {
		fmt.Println("Failed to get order cuz", err)
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var details string
		var id int
		err = rows.Scan(&id, &details)

		if err != nil {
			fmt.Println("Failed to populate order cuz", err)
			return orders, err
		}

		orders[id] = details
	}

	return orders, nil
}

func GetAllOrders(conn *sql.DB) (map[int]string, error) {
	rows, err := conn.Query("select id, details from orders")
	orders := make(map[int]string)

	if err != nil {
		fmt.Println("Failed to get order cuz", err)
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var details string
		var id int
		err = rows.Scan(&id, &details)

		if err != nil {
			fmt.Println("Failed to populate order cuz", err)
			return orders, err
		}

		orders[id] = details
	}

	return orders, nil
}

func DeleteOrders(conn *sql.DB, IDs []int) error {
	stmt := "delete from orders where id in ( "

	for _, id := range IDs {
		stmt += fmt.Sprintf("%d, ", id)
	}

	stmt = stmt[:len(stmt)-2]

	_, err := conn.Exec(stmt + " )")

	if err != nil {
		fmt.Println("Failed to delete order cuz", err)
		return err
	}

	stmt = "delete from users_orders_mapping where order_id in ( "

	for _, id := range IDs {
		stmt += fmt.Sprintf("%d, ", id)
	}

	stmt = stmt[:len(stmt)-2]

	_, err = conn.Exec(stmt + " )")

	return nil
}

func DeleteAllOrders(conn *sql.DB) error {
	_, err := conn.Exec("delete from orders")

	if err != nil {
		fmt.Println("Failed to delete order cuz", err)
		return err
	}

	_, err = conn.Exec("delete from users_orders_mapping")

	if err != nil {
		fmt.Println("Failed to delete order cuz", err)
		return err
	}

	return nil
}
