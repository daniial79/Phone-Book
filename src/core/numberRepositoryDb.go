package core

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"strconv"
)

// NumberRepositoryDb secondary actor
type NumberRepositoryDb struct {
	client *sql.DB
}

func NewNumberRepositoryDb(client *sql.DB) NumberRepository {
	return NumberRepositoryDb{client}
}

func (r NumberRepositoryDb) AddNewNumber(n Number) (*Number, *errs.AppError) {
	insertSql := `INSERT INTO numbers(contact_id, number, label) VALUES($1, $2, $3) RETURNING id`

	var integerId int
	row := r.client.QueryRow(insertSql, n.ContactId, n.PhoneNumber, n.Label)
	err := row.Scan(&integerId)
	if err != nil {
		logger.Error("Error while retrieving id for last inserted number into existing contact: " + err.Error())
		return nil, errs.NewUnexpectedErr("Unexpected error happened")
	}

	lastInsertedId := strconv.Itoa(integerId)
	n.Id = lastInsertedId

	return &n, nil
}
