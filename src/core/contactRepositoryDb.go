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

func NewContactRepositoryDb(client *sql.DB) ContactRepositoryDb {
	return ContactRepositoryDb{client: client}
}

func (r ContactRepositoryDb) CreateContact(c *Contact) (*Contact, *errs.AppError) {
	//TODO: Find better approach with smaller Space Complexity
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
