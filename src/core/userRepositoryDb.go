package core

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/lib/pq"
)

type UserRepositoryDb struct {
	client *sql.DB
}

func NewUserRepositoryDb(client *sql.DB) UserRepositoryDb {
	return UserRepositoryDb{client: client}
}

func (r UserRepositoryDb) CreateUser(u User) (*User, *errs.AppError) {
	insertSql := `INSERT INTO users(username, password, phone_number, created_at, updated_at) 
				  VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var insertedId string
	row := r.client.QueryRow(insertSql, u.Username, u.Password, u.PhoneNumber, u.CreatedAt, u.UpdatedAt)

	if err := row.Scan(&insertedId); err != nil {
		fmt.Println(err)

		var pgerr *pq.Error
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				return nil, errs.NewUnProcessableErr(errs.UsernameUniquenessViolationErr)
			}
		}

		logger.Error("Error while inserting new record to user table")
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	u.Id = insertedId
	return &u, nil
}

func (r UserRepositoryDb) GetUserByUsername(username string) (*User, *errs.AppError) {
	selectQuery := `SELECT id, username, password FROM users WHERE username = $1`

	var fetchedUser User
	row := r.client.QueryRow(selectQuery, username)
	err := row.Scan(&fetchedUser.Id, &fetchedUser.Username, &fetchedUser.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundErr("There is no user with such username")
		}
		logger.Error("Error while getting user by username: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	return &fetchedUser, nil
}
