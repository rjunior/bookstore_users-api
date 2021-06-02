package users

import (
	"github.com/rjunior/bookstore_users-api/datasources/mysql/users_db"
	"github.com/rjunior/bookstore_users-api/utils/date_utils"
	"github.com/rjunior/bookstore_users-api/utils/errors"
	mysqlutils "github.com/rjunior/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			panic(errClose.Error())
		}
	}()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			panic(errClose.Error())
		}
	}()

	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			panic(errClose.Error())
		}
	}()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}
