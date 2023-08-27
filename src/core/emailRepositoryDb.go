package core

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"strconv"
)

// EmailRepositoryDb secondary actor
type EmailRepositoryDb struct {
	client *sql.DB
}

func (r EmailRepositoryDb) Create(e Email) (*Email, *errs.AppError) {
	insertSql := `INSERT INTO emails(contact_id, address) VALUES ($1, $2) RETURNING id`

	result, err := r.client.Exec(insertSql, e.ContactId, e.Address)
	if err != nil {
		logger.Error("Error while inserting record to emails table: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected internal error")
	}

	integerId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving last inserted record to emails table to string: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected internal error")
	}

	strId := strconv.Itoa(int(integerId))
	e.ID = strId

	return &e, nil

}
