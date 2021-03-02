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

func UpdateUser(conn *sql.DB, newName string) {

}

func DeleteUser(conn *sql.DB, id int) {

}
