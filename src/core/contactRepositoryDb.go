package core

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"strconv"
)

// ContactRepositoryDb Secondary actor
type ContactRepositoryDb struct {
	client *sql.DB
}

func (r ContactRepositoryDb) CreateEmail(e Email) (*Email, *errs.AppError) {
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

func (r ContactRepositoryDb) CreateNumber(n Number) (*Number, *errs.AppError) {
	insertSql := `INSERT INTO numbers(id, contact_id, number, label) VALUES ($1, $2) RETURNING id`

	result, err := r.client.Exec(insertSql, n.ContactId, n.Number, n.Label)
	if err != nil {
		logger.Error("Error while inserting record to numbers table: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected internal error")
	}

	integerId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving last inserted record to numbers table to string: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected internal error")
	}

	strId := strconv.Itoa(int(integerId))
	n.ID = strId

	return &n, nil

}

func (r ContactRepositoryDb) CreateContact(c Contact) (*Contact, *errs.AppError) {
	insertSql := `INSERT INTO contacts(first_name, last_name) VALUES ($1, $2) RETURNING id`

	result, err := r.client.Exec(insertSql, c.FirstName, c.LastName)
	if err != nil {
		logger.Error("Error while inserting record to contacts table: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected internal error")
	}

	integerId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving last inserted record to contacts table to string: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected internal error")
	}

	strId := strconv.Itoa(int(integerId))
	c.Id = strId

	return &c, nil

}
