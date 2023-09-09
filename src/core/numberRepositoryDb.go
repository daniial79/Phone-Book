package core

import (
	"database/sql"
	"errors"
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

func (r NumberRepositoryDb) CheckContactExistenceById(cId string) *errs.AppError {
	var contactId int
	checkContactSql := `SELECT id FROM contacts WHERE id =  $1`
	row := r.client.QueryRow(checkContactSql, cId)
	err := row.Scan(&contactId)
	if errors.Is(err, sql.ErrNoRows) {
		return errs.NewNotFoundErr("contact with such id is not found")
	}
	return nil
}

func (r NumberRepositoryDb) AddNewNumber(n Number) (*Number, *errs.AppError) {
	//checking the existence of contact with specified id
	appErr := r.CheckContactExistenceById(n.ContactId)
	if appErr != nil {
		return nil, appErr
	}

	//adding number associated with existing contact id
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
