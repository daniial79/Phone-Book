package core

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/google/uuid"
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
		logger.Error("Error while inserting new record to user table")
		return nil, errs.NewUnexpectedErr("Unexpected error happened")
	}

	u.Id = insertedId
	return &u, nil
}
