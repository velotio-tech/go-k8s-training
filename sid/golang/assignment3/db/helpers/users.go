package helpers

import (
	"database/sql"
	"fmt"
)

func CreateUser(conn *sql.DB, name string) error {
	_, err := conn.Exec(fmt.Sprintf("insert into users (name) values ('%s')", name))

	if err != nil {
		fmt.Println("Failed to create user cuz", err)
		return err
	}

	return nil
}

func GetUser(conn *sql.DB, id int) (map[int]string, error) {
	rows, err := conn.Query(fmt.Sprintf("select id, name from users where id=%d", id))
	users := make(map[int]string)

	if err != nil {
		fmt.Println("Failed to get user cuz", err)
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var id int
		err = rows.Scan(&id, &name)

		if err != nil {
			fmt.Println("Failed to populate user cuz", err)
			return users, err
		}

		users[id] = name
	}

	return users, nil
}

func UpdateUser(conn *sql.DB, id int, newName string) error {
	_, err := conn.Exec(fmt.Sprintf("update users set name = '%s' where id = %d", newName, id))

	if err != nil {
		fmt.Println("Failed to update user cuz", err)
		return err
	}

	return nil
}

func DeleteUser(conn *sql.DB, id int) error {
	_, err := conn.Exec(fmt.Sprintf("delete from users where id = %d", id))

	if err != nil {
		fmt.Println("Failed to delete user cuz", err)
		return err
	}

	_, err = conn.Exec(fmt.Sprintf("delete from users_orders_mapping where user_id = %d", id))

	if err != nil {
		fmt.Println("Failed to delete user order mapping cuz", err)
		return err
	}

	return nil
}

func GetUsers(conn *sql.DB, IDs []int) (map[int]string, error) {
	stmt := "select id, name from users where id in ( "

	for _, id := range IDs {
		stmt += fmt.Sprintf("%d, ", id)
	}

	stmt = stmt[:len(stmt)-2]

	rows, err := conn.Query(stmt + " )")
	users := make(map[int]string)

	if err != nil {
		fmt.Println("Failed to get user cuz", err)
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var details string
		var id int
		err = rows.Scan(&id, &details)

		if err != nil {
			fmt.Println("Failed to populate user cuz", err)
			return users, err
		}

		users[id] = details
	}

	return users, nil
}

func GetAllUsers(conn *sql.DB) (map[int]string, error) {
	rows, err := conn.Query("select id, name from users")
	users := make(map[int]string)

	if err != nil {
		fmt.Println("Failed to get user cuz", err)
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var details string
		var id int
		err = rows.Scan(&id, &details)

		if err != nil {
			fmt.Println("Failed to populate user cuz", err)
			return users, err
		}

		users[id] = details
	}

	return users, nil
}

func DeleteUsers(conn *sql.DB, IDs []int) error {
	stmt := "delete from users where id in ( "

	for _, id := range IDs {
		stmt += fmt.Sprintf("%d, ", id)
	}

	stmt = stmt[:len(stmt)-2]

	_, err := conn.Exec(stmt + " )")

	if err != nil {
		fmt.Println("Failed to delete user cuz", err)
		return err
	}

	stmt = "delete from users_users_mapping where user_id in ( "

	for _, id := range IDs {
		stmt += fmt.Sprintf("%d, ", id)
	}

	stmt = stmt[:len(stmt)-2]

	_, err = conn.Exec(stmt + " )")

	return nil
}

func DeleteAllUsers(conn *sql.DB) error {
	_, err := conn.Exec("delete from users")

	if err != nil {
		fmt.Println("Failed to delete user cuz", err)
		return err
	}

	_, err = conn.Exec("delete from users_users_mapping")

	if err != nil {
		fmt.Println("Failed to delete user cuz", err)
		return err
	}

	return nil
}
