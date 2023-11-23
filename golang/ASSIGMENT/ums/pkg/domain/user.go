package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"ums/pkg/exception"
	"ums/pkg/models"
)

// CreateUser create a new user
func (u *UserCliet) CreateUser(user *models.User) (string, *exception.Exception) {
	result, err := u.DB.Exec(`INSERT IGNORE INTO users (name,email) VALUES ($1,$2) RETURNING "ID"`, user.Name, user.Email)
	if err != nil {
		log.Println("failed to create user : ", err)
		return "", &exception.Exception{
			Err:        ErrUserCreationFailed,
			Message:    ErrUserCreationFailed.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_DB_ERROR,
		}
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println("failed to retirve rows affected data : ", err)
		return "", &exception.Exception{
			Err:        ErrUserCreationFailed,
			Message:    ErrUserCreationFailed.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_DB_ERROR,
		}
	} else if rowsAffected == 0 {
		return "", &exception.Exception{
			Err:        ErrUserAlreadyExists,
			Message:    ErrUserAlreadyExists.Error(),
			StatusCode: http.StatusConflict,
			StatusText: exception.STATUS_DB_ERROR,
		}
	}

	return fmt.Sprint(result.LastInsertId()), nil
}

// GetUser will return only the user details on the basis of either ID whihc is the primary key of the user or by the email id.
// Note :  If both the details are passed, we are going to only consider the primary key
func (u *UserCliet) GetUser(ID, email string) (*models.User, *exception.Exception) {
	var getUserQuery string
	var err error
	user := &models.User{}

	if len(ID) > 0 {
		getUserQuery = `SELECT "ID","email","createdAt","updatedAt" FROM "users" WHERE "ID"=$1'`
		err = u.DB.QueryRow(getUserQuery, ID).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	} else {
		getUserQuery = `SELECT "ID","email","createdAt","updatedAt" FROM "users" WHERE "email"=$1'`
		err = u.DB.QueryRow(getUserQuery, email).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	}

	if err != nil {
		log.Println("failed to retrive the user : ", err)
		if errors.Is(sql.ErrNoRows, err) {
			return nil, &exception.Exception{
				Err:        ErrUserNotFound,
				Message:    ErrUserNotFound.Error(),
				StatusCode: http.StatusNotFound,
				StatusText: exception.STATUS_DB_ERROR,
			}
		}
		return nil, &exception.Exception{
			Err:        ErrInternalError,
			Message:    ErrInternalError.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_DB_ERROR,
		}
	}

	return user, nil
}

// GetUsers get all the users in the databases
func (u *UserCliet) GetUsers() ([]models.User, *exception.Exception) {
	users := []models.User{}

	getUsersQuery := `SELECT "ID","email","createdAt","updatedAt" FROM "users";`
	rows, err := u.DB.Query(getUsersQuery)
	if err != nil {
		log.Println("failed to fetch users : ", err)
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil
		}
		return nil, &exception.Exception{
			Err:        ErrInternalError,
			Message:    ErrInternalError.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_DB_ERROR,
		}
	}

	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates a user in the database
func (u *UserCliet) UpdateUser(ID, name string) *exception.Exception {
	updateUserQuery := `UPDATE "users" SET "name"=$1 WHERE "ID"=$2'`
	result, err := u.DB.Exec(updateUserQuery, name, ID)
	if err != nil {
		log.Println("failed to update the user : ", err)
		return &exception.Exception{
			Err:        ErrFailedToUpdateUser,
			Message:    ErrFailedToUpdateUser.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_DB_ERROR,
		}
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return &exception.Exception{
			Err:        ErrFailedToUpdateUser,
			Message:    ErrFailedToUpdateUser.Error(),
			StatusCode: http.StatusInternalServerError,
			StatusText: exception.STATUS_DB_ERROR,
		}
	}

	return nil
}
