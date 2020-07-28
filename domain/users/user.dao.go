package users

import (
	"fmt"
	"github.com/alvarezcarlos/bookstore_users-api/datasources/mysql/users_db"
	"github.com/alvarezcarlos/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	indexUniqueEmail      = "email_UNIQUE"
	errorNoRows           = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	querySearchUser       = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(querySearchUser)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.InternalServerError(err.Error())
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	//insertResult, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		sqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return errors.InternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
		}

		switch sqlError.Number {
		case 1062:
			return errors.InternalServerError(fmt.Sprintf("email %s already exists", user.Email))
		}
		if strings.Contains(err.Error(), indexUniqueEmail) {

		}
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}

	user.Id = userId

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.InternalServerError("unable to retrieve users")
		//return []User{}, errors.InternalServerError("unable to retrieve users")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.InternalServerError("unable to retrieve users")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user *User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.InternalServerError(fmt.Sprintf(err.Error()))
		}
		results = append(results, *user)
	}
	return results, nil
}
