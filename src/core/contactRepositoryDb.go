package core

import (
	"database/sql"
	"errors"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"strconv"
)

// ContactRepositoryDb Secondary actor
type ContactRepositoryDb struct {
	client *sql.DB
}

func NewContactRepositoryDb(client *sql.DB) ContactRepositoryDb {
	return ContactRepositoryDb{client: client}
}

func (r ContactRepositoryDb) CreateContact(c *Contact) (*Contact, *errs.AppError) {
	insertedNumbers := make([]Number, 0)
	insertedEmails := make([]Email, 0)
	tx, err := r.client.Begin()

	//inserting new record to contact tables
	if err != nil {
		logger.Error("Error while starting transaction in order to create new contact: " + err.Error())
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	contactInsertSql := `INSERT INTO contacts(first_name, last_name) VALUES($1, $2) RETURNING id`
	cRow := tx.QueryRow(contactInsertSql, c.FirstName, c.LastName)

	var cIntegerId int
	err = cRow.Scan(&cIntegerId)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Error("Error while rollback the crate contact transaction: " + txErr.Error())
		}
		logger.Error("Error while fetching last inserted id from contact tables: " + err.Error())
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	cStringId := strconv.Itoa(cIntegerId)
	c.Id = cStringId

	//inserting new record(s) into numbers
	insertNumbersSql := `INSERT INTO numbers(contact_id, number, label) VALUES ($1, $2, $3) RETURNING id`
	for _, number := range c.PhoneNumbers {
		nRow := tx.QueryRow(insertNumbersSql,
			c.Id,
			number.PhoneNumber,
			number.Label,
		)

		var nIntegerId int
		err = nRow.Scan(&nIntegerId)

		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				logger.Error("Error while rollback the crate contact transaction: " + txErr.Error())
			}
			logger.Error("Error while fetching last inserted id from numbers tables: " + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected database error")
		}

		nStringId := strconv.Itoa(nIntegerId)
		number.Id = nStringId
		number.ContactId = c.Id
		insertedNumbers = append(insertedNumbers, number)
	}

	//inserting new record into emails
	insertEmailSql := `INSERT INTO emails(contact_id, address) VALUES ($1, $2) RETURNING id`
	for _, email := range c.Emails {
		eRow := tx.QueryRow(insertEmailSql, c.Id, email.Address)

		var eIntegerId int
		err = eRow.Scan(&eIntegerId)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				logger.Error("Error while rollback the crate contact transaction: " + txErr.Error())
			}
			logger.Error("Error while fetching last inserted id from emails tables: " + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected database error")
		}

		eStringId := strconv.Itoa(eIntegerId)
		email.Id = eStringId
		email.ContactId = c.Id

		insertedEmails = append(insertedEmails, email)
	}

	txErr := tx.Commit()
	if txErr != nil {
		logger.Error("Error while committing the new created contact transaction")
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	c.PhoneNumbers = insertedNumbers
	c.Emails = insertedEmails

	return c, nil
}

func (r ContactRepositoryDb) CheckContactExistenceById(cId string) *errs.AppError {
	var contactId int
	checkContactSql := `SELECT id FROM contacts WHERE id =  $1`
	row := r.client.QueryRow(checkContactSql, cId)
	err := row.Scan(&contactId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr("contact with such id is not found")
		}
		logger.Error("Error while retrieving contact id for existence check (number repo): " + err.Error())
		return errs.NewUnexpectedErr("Unexpected error happened")
	}
	return nil
}

func (r ContactRepositoryDb) AddNewNumber(n []Number) ([]Number, *errs.AppError) {
	//checking the existence of contact with specified id
	for _, number := range n {
		appErr := r.CheckContactExistenceById(number.ContactId)
		if appErr != nil {
			return nil, appErr
		}
	}

	//adding number associated with existing contact id
	result := make([]Number, len(n))
	insertSql := `INSERT INTO numbers(contact_id, number, label) VALUES($1, $2, $3) RETURNING id`
	for i, number := range n {

		var integerId int
		row := r.client.QueryRow(insertSql, number.ContactId, number.PhoneNumber, number.Label)
		err := row.Scan(&integerId)
		if err != nil {
			logger.Error("Error while retrieving id for last inserted number into existing contact: " + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected error happened")

		}
		lastInsertedId := strconv.Itoa(integerId)
		number.Id = lastInsertedId

		result[i] = number

	}

	return result, nil
}

func (r ContactRepositoryDb) AddNewEmails(e []Email) ([]Email, *errs.AppError) {
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

func (r ContactRepositoryDb) GetAllContacts(filters map[string]string) ([]Contact, *errs.AppError) {
	//TODO: find a fucking solution for joining tables :/
	return nil, errs.NewUnexpectedErr("Unexpected error happened")
}
