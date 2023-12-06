package core

import (
	"database/sql"
	"errors"
	"fmt"
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
			return nil, errs.NewUnexpectedErr("Unexpected error happened")
		}

		lastInsertedId := strconv.Itoa(integerId)
		email.Id = lastInsertedId

		result[i] = email
	}

	return result, nil
}

func (r ContactRepositoryDb) GetAllContacts() ([]Contact, *errs.AppError) {
	contacts := make([]Contact, 0)
	selectContactSql := `SELECT * FROM contacts`

	rows, err := r.client.Query(selectContactSql)
	if err != nil {
		logger.Error("Error while querying contacts table: " + err.Error())
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.Id, &c.FirstName, &c.LastName)
		if err != nil {
			logger.Error("Error while scanning retrieved records from contacts table: " + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected database error")
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (r ContactRepositoryDb) GetContactNumbers(cId string) ([]Number, *errs.AppError) {
	numbers := make([]Number, 0)

	selectNumbersSql := `SELECT * FROM numbers WHERE contact_id = $1`

	rows, err := r.client.Query(selectNumbersSql, cId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundErr("no number associated with this contact id")
		}
		logger.Error("Error while selecting numbers associated with specific contact: " + err.Error())
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	for rows.Next() {
		var n Number
		err = rows.Scan(&n.Id, &n.ContactId, &n.PhoneNumber, &n.Label)
		if err != nil {
			logger.Error("Error while scanning retrieved set of numbers: " + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected database error")
		}
		numbers = append(numbers, n)
	}

	return numbers, nil

}

func (r ContactRepositoryDb) GetContactEmails(cId string) ([]Email, *errs.AppError) {
	emails := make([]Email, 0)

	selectEmailsSql := `SELECT * FROM emails WHERE contact_id = $1`
	rows, err := r.client.Query(selectEmailsSql, cId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundErr("no email associated with this contact id")
		}
		logger.Error("Error while selecting emails associated with specific contact: " + err.Error())
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	for rows.Next() {
		var e Email
		err = rows.Scan(&e.Id, &e.ContactId, &e.Address)
		if err != nil {
			logger.Error("Error while scanning retrieved set of emails: " + err.Error())
			return nil, errs.NewUnexpectedErr("Unexpected database error")
		}
		emails = append(emails, e)
	}

	return emails, nil
}

func (r ContactRepositoryDb) DeleteContactEmail(cId, eId string) *errs.AppError {
	deleteQuery := `DELETE FROM emails WHERE id = $1 AND contact_id = $2 RETURNING id`

	row := r.client.QueryRow(deleteQuery, eId, cId)
	var deletedEmailId int

	err := row.Scan(&deletedEmailId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr("email not found")
		}
		logger.Error("Error while removing a record from emails table: " + err.Error())
		return errs.NewUnexpectedErr("Unexpected database error")
	}

	return nil
}

func (r ContactRepositoryDb) DeleteContactPhoneNumber(cId, nId string) *errs.AppError {
	deleteQuery := `DELETE FROM numbers WHERE id = $1 AND contact_id = $2 RETURNING id`

	row := r.client.QueryRow(deleteQuery, nId, cId)
	var deletedNumberId int

	err := row.Scan(&deletedNumberId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr("phone number not found")
		}
		return errs.NewUnexpectedErr("Unexpected database error")
	}

	return nil
}

func (r ContactRepositoryDb) DeleteContact(cId string) *errs.AppError {
	tx, err := r.client.Begin()
	if err != nil {
		logger.Error("Error while starting new transaction for cascading deletion of contact: " + err.Error())
		return errs.NewUnexpectedErr("Unexpected database error")
	}

	pnDeleteQuery := `DELETE FROM numbers WHERE contact_id = $1`
	_, err = tx.Exec(pnDeleteQuery, cId)

	if err != nil {
		logger.Error("Error while removing corresponding phone numbers with contact from numbers table: " + err.Error())
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Error("Error while rollback from cascading deletion on numbers table:" + txErr.Error())
			return errs.NewUnexpectedErr("Unexpected database error")
		}
	}

	eDeleteQuery := `DELETE FROM emails WHERE contact_id = $1`
	_, err = tx.Exec(eDeleteQuery, cId)

	if err != nil {
		logger.Error("Error while removing corresponding emails with contact from emails table: " + err.Error())
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Error("Error while rollback from cascading deletion on emails table:" + txErr.Error())
			return errs.NewUnexpectedErr("Unexpected database error")
		}
	}

	cDeleteQuery := `DELETE FROM contacts WHERE id = $1 RETURNING id`
	row := tx.QueryRow(cDeleteQuery, cId)
	var deletedContactId int

	if err = row.Scan(&deletedContactId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr("contact not found")
		}

		if txErr := tx.Rollback(); txErr != nil {
			logger.Error("Error while rollback from deleting record from contacts table:" + txErr.Error())
			return errs.NewUnexpectedErr("Unexpected database error")
		}
	}

	if txErr := tx.Commit(); txErr != nil {
		logger.Error("Error while committing the started transaction inside contact deletion")
	}

	return nil
}

func (r ContactRepositoryDb) UpdateContactPhoneNumber(newNumber Number) (*Number, *errs.AppError) {
	numberId := newNumber.Id
	contactId := newNumber.ContactId
	newPhoneNumber := newNumber.PhoneNumber
	newLabel := newNumber.Label

	var row *sql.Row
	if newPhoneNumber != "" && newLabel != "" {
		fmt.Println("both")
		updateQuery := `UPDATE numbers SET number = $1, label = $2 WHERE id = $3 AND contact_id = $4 RETURNING id`
		row = r.client.QueryRow(updateQuery, newPhoneNumber, newLabel, numberId, contactId)
	} else if newPhoneNumber == "" && newLabel != "" {
		fmt.Println("label")
		updateQuery := `UPDATE numbers SET label = $1 WHERE id = $2 AND contact_id = $3 RETURNING id`
		row = r.client.QueryRow(updateQuery, newLabel, numberId, contactId)
	} else if newPhoneNumber != "" && newLabel == "" {
		fmt.Println("number")
		updateQuery := `UPDATE numbers SET number = $1 WHERE id = $2 AND contact_id = $3 RETURNING id`
		row = r.client.QueryRow(updateQuery, newPhoneNumber, numberId, contactId)
	}

	var retrievedId int
	err := row.Scan(&retrievedId)
	if err != nil {
		logger.Error("Error while updating record in numbers table")
		return nil, errs.NewUnexpectedErr("Unexpected database error")
	}

	return &newNumber, nil
}
