package core

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/google/uuid"
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

	var insertedId uuid.UUID
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
