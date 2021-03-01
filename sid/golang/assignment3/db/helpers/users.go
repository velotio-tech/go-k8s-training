package helpers

import (
	"database/sql"
	"fmt"
)

func CreateUser(conn *sql.DB, name string) error {
	stmt := fmt.Sprintf("insert into users (name) values ('%s')", name)

	fmt.Println(stmt)

	_, err := conn.Exec(stmt)

	if err != nil {
		fmt.Println("Failed to create user cuz", err)
		return err
	}

	return nil
}

func GetUser(conn *sql.DB, id int) {

}

func UpdateUser(conn *sql.DB, newName string) {

}

func DeleteUser(conn *sql.DB, id int) {

}
