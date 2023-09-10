package core

import (
	"database/sql"
	"errors"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"strconv"
)

// EmailRepositoryDb secondary actor
type EmailRepositoryDb struct {
	client *sql.DB
}

func NewEmailRepositoryDb(client *sql.DB) EmailRepositoryDb {
	return EmailRepositoryDb{client}
}

func (r EmailRepositoryDb) CheckContactExistenceById(cId string) *errs.AppError {
	var contactId int
	checkContactSql := `SELECT id FROM contacts WHERE id = $1`
	row := r.client.QueryRow(checkContactSql, cId)
	err := row.Scan(&contactId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr("contact with such id not found")
		}
		logger.Error("Error while retrieving contact id for existence check (email repo): " + err.Error())
		return errs.NewUnexpectedErr("Unexpected error happened")
	}
	return nil
}

func (r EmailRepositoryDb) AddNewEmails(e []Email) ([]Email, *errs.AppError) {
	for _, email := range e {
		appErr := r.CheckContactExistenceById(email.ContactId)
		if appErr != nil {
			return nil, appErr
		}
	}

	result := make([]Email, len(e))
	insertSql := `INSERT INTO emails(contact_id, address) VALUES($1, $2) RETURNING id`
	for i, email := range e {
		var integerId int
		row := r.client.QueryRow(insertSql, email.ContactId, email.Address)
		err := row.Scan(&integerId)
		if err != nil {
			logger.Error("Error while retrieving id for last inserted email into existing contact:" + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected error happend")
		}

		lastInsertedId := strconv.Itoa(integerId)
		email.Id = lastInsertedId

		result[i] = email
	}

	return result, nil
}
