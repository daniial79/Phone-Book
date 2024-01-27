package core

import (
	"database/sql"
	"errors"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/google/uuid"
)

// ContactRepositoryDb Secondary actor
type ContactRepositoryDb struct {
	client *sql.DB
}

func NewContactRepositoryDb(client *sql.DB) ContactRepositoryDb {
	return ContactRepositoryDb{client: client}
}

// GetContactOwnerByUsername Check for user existence and user id eventually using username
func (r ContactRepositoryDb) GetContactOwnerByUsername(username string) (uuid.UUID, *errs.AppError) {
	selectQuery := `SELECT id FROM users WHERE username = $1`

	var userId uuid.UUID
	row := r.client.QueryRow(selectQuery, username)
	err := row.Scan(&userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.UUID{}, errs.NewNotFoundErr(errs.UserNotFoundErr)
		}

		return uuid.UUID{}, errs.NewUnexpectedErr(errs.InternalErr)
	}

	return userId, nil
}

// CreateContact Initiate new contact if owner of contact exists
func (r ContactRepositoryDb) CreateContact(username string, c *Contact) (*Contact, *errs.AppError) {

	// find contact owner
	ownerId, appErr := r.GetContactOwnerByUsername(username)
	if appErr != nil {
		return nil, appErr
	}
	c.OwnerId = ownerId

	insertedNumbers := make([]Number, 0)
	insertedEmails := make([]Email, 0)
	tx, err := r.client.Begin()

	//inserting new record to contact tables
	if err != nil {
		logger.Error("Error while starting transaction in order to create new contact: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	contactInsertSql := `INSERT INTO contacts(first_name, last_name) VALUES($1, $2) WHERE owner_id = $3 RETURNING id`
	cRow := tx.QueryRow(contactInsertSql, c.FirstName, c.LastName, c.OwnerId)

	var insertedContactId uuid.UUID
	err = cRow.Scan(&insertedContactId)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Error("Error while rollback the crate contact transaction: " + txErr.Error())
		}
		logger.Error("Error while fetching last inserted id from contact tables: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	// This will be used for creating contact assets
	c.Id = insertedContactId

	//inserting new record(s) into numbers
	insertNumbersSql := `INSERT INTO numbers(contact_id, number, label) VALUES ($1, $2, $3) RETURNING id`
	for _, number := range c.PhoneNumbers {
		nRow := tx.QueryRow(insertNumbersSql,
			c.Id,
			number.PhoneNumber,
			number.Label,
		)

		var insertedNumberId uuid.UUID
		err = nRow.Scan(&insertedNumberId)

		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				logger.Error("Error while rollback the crate contact transaction: " + txErr.Error())
			}
			logger.Error("Error while fetching last inserted id from numbers tables: " + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)
		}

		number.Id = insertedNumberId
		number.ContactId = c.Id
		insertedNumbers = append(insertedNumbers, number)
	}

	//inserting new record into emails
	insertEmailSql := `INSERT INTO emails(contact_id, address) VALUES ($1, $2) RETURNING id`
	for _, email := range c.Emails {
		eRow := tx.QueryRow(insertEmailSql, c.Id, email.Address)

		var insertedEmailId uuid.UUID
		err = eRow.Scan(&insertedEmailId)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				logger.Error("Error while rollback the crate contact transaction: " + txErr.Error())
			}
			logger.Error("Error while fetching last inserted id from emails tables: " + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)
		}

		email.Id = insertedEmailId
		email.ContactId = c.Id

		insertedEmails = append(insertedEmails, email)
	}

	txErr := tx.Commit()
	if txErr != nil {
		logger.Error("Error while committing the new created contact transaction")
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	c.PhoneNumbers = insertedNumbers
	c.Emails = insertedEmails

	return c, nil
}

func (r ContactRepositoryDb) CheckContactExistenceById(cId uuid.UUID) *errs.AppError {
	var contactId uuid.UUID
	checkContactSql := `SELECT id FROM contacts WHERE id =  $1`
	row := r.client.QueryRow(checkContactSql, cId)
	err := row.Scan(&contactId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr(errs.ContactNotFoundErr)
		}
		logger.Error("Error while retrieving contact id for existence check (number repo): " + err.Error())
		return errs.NewUnexpectedErr(errs.InternalErr)
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

		var insertedId uuid.UUID
		row := r.client.QueryRow(insertSql, number.ContactId, number.PhoneNumber, number.Label)
		err := row.Scan(&insertedId)
		if err != nil {
			logger.Error("Error while retrieving id for last inserted number into existing contact: " + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)

		}

		number.Id = insertedId

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
		var insertedRecordId uuid.UUID
		row := r.client.QueryRow(insertSql, email.ContactId, email.Address)
		err := row.Scan(&insertedRecordId)
		if err != nil {
			logger.Error("Error while retrieving id for last inserted email into existing contact:" + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)
		}

		email.Id = insertedRecordId

		result[i] = email
	}

	return result, nil
}

func (r ContactRepositoryDb) GetAllContacts(username string) ([]Contact, *errs.AppError) {
	ownerId, appErr := r.GetContactOwnerByUsername(username)
	if appErr != nil {
		return nil, errs.NewNotFoundErr(errs.UserNotFoundErr)
	}

	contacts := make([]Contact, 0)
	selectContactSql := `SELECT * FROM contacts WHERE owner_id = $1`

	rows, err := r.client.Query(selectContactSql, ownerId)
	if err != nil {
		logger.Error("Error while querying contacts table: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.Id, &c.FirstName, &c.LastName)
		if err != nil {
			logger.Error("Error while scanning retrieved records from contacts table: " + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (r ContactRepositoryDb) GetContactNumbers(cId uuid.UUID) ([]Number, *errs.AppError) {
	numbers := make([]Number, 0)

	selectNumbersSql := `SELECT * FROM numbers WHERE contact_id = $1`

	rows, err := r.client.Query(selectNumbersSql, cId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundErr("no number associated with this contact id")
		}
		logger.Error("Error while selecting numbers associated with specific contact: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	for rows.Next() {
		var n Number
		err = rows.Scan(&n.Id, &n.ContactId, &n.PhoneNumber, &n.Label)
		if err != nil {
			logger.Error("Error while scanning retrieved set of numbers: " + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)
		}
		numbers = append(numbers, n)
	}

	return numbers, nil

}

func (r ContactRepositoryDb) GetContactEmails(cId uuid.UUID) ([]Email, *errs.AppError) {
	emails := make([]Email, 0)

	selectEmailsSql := `SELECT * FROM emails WHERE contact_id = $1`
	rows, err := r.client.Query(selectEmailsSql, cId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundErr("no email associated with this contact id")
		}
		logger.Error("Error while selecting emails associated with specific contact: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	for rows.Next() {
		var e Email
		err = rows.Scan(&e.Id, &e.ContactId, &e.Address)
		if err != nil {
			logger.Error("Error while scanning retrieved set of emails: " + err.Error())
			return nil, errs.NewUnexpectedErr(errs.InternalErr)
		}
		emails = append(emails, e)
	}

	return emails, nil
}

func (r ContactRepositoryDb) DeleteContactEmail(cId, eId uuid.UUID) *errs.AppError {
	deleteQuery := `DELETE FROM emails WHERE id = $1 AND contact_id = $2 RETURNING id`

	row := r.client.QueryRow(deleteQuery, eId, cId)
	var deletedEmailId int

	err := row.Scan(&deletedEmailId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr(errs.EmailNotFoundErr)
		}
		logger.Error("Error while removing a record from emails table: " + err.Error())
		return errs.NewUnexpectedErr(errs.InternalErr)
	}

	return nil
}

func (r ContactRepositoryDb) DeleteContactPhoneNumber(cId, nId uuid.UUID) *errs.AppError {
	deleteQuery := `DELETE FROM numbers WHERE id = $1 AND contact_id = $2 RETURNING id`

	row := r.client.QueryRow(deleteQuery, nId, cId)
	var deletedNumberId int

	err := row.Scan(&deletedNumberId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr(errs.NumberNotFoundErr)
		}
		return errs.NewUnexpectedErr(errs.InternalErr)
	}

	return nil
}

func (r ContactRepositoryDb) DeleteContact(cId uuid.UUID) *errs.AppError {

	tx, err := r.client.Begin()
	if err != nil {
		logger.Error("Error while starting new transaction for cascading deletion of contact: " + err.Error())
		return errs.NewUnexpectedErr(errs.InternalErr)
	}

	pnDeleteQuery := `DELETE FROM numbers WHERE contact_id = $1`
	_, err = tx.Exec(pnDeleteQuery, cId)

	if err != nil {
		logger.Error("Error while removing corresponding phone numbers with contact from numbers table: " + err.Error())
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Error("Error while rollback from cascading deletion on numbers table:" + txErr.Error())
			return errs.NewUnexpectedErr(errs.InternalErr)
		}
	}

	eDeleteQuery := `DELETE FROM emails WHERE contact_id = $1`
	_, err = tx.Exec(eDeleteQuery, cId)

	if err != nil {
		logger.Error("Error while removing corresponding emails with contact from emails table: " + err.Error())
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Error("Error while rollback from cascading deletion on emails table:" + txErr.Error())
			return errs.NewUnexpectedErr(errs.InternalErr)
		}
	}

	cDeleteQuery := `DELETE FROM contacts WHERE id = $1 AND owner_id = $2 RETURNING id`
	row := tx.QueryRow(cDeleteQuery, cId)
	var deletedContactId int

	if err = row.Scan(&deletedContactId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundErr("contact not found")
		}

		if txErr := tx.Rollback(); txErr != nil {
			logger.Error("Error while rollback from deleting record from contacts table:" + txErr.Error())
			return errs.NewUnexpectedErr(errs.InternalErr)
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

	if appErr := r.CheckContactExistenceById(contactId); appErr != nil {
		return nil, appErr
	}

	var row *sql.Row
	if newPhoneNumber != "" && newLabel != "" {
		updateQuery := `UPDATE numbers SET number = $1, label = $2 WHERE id = $3 AND contact_id = $4 RETURNING id`
		row = r.client.QueryRow(updateQuery, newPhoneNumber, newLabel, numberId, contactId)
	} else if newPhoneNumber == "" && newLabel != "" {
		updateQuery := `UPDATE numbers SET label = $1 WHERE id = $2 AND contact_id = $3 RETURNING id`
		row = r.client.QueryRow(updateQuery, newLabel, numberId, contactId)
	} else if newPhoneNumber != "" && newLabel == "" {
		updateQuery := `UPDATE numbers SET number = $1 WHERE id = $2 AND contact_id = $3 RETURNING id`
		row = r.client.QueryRow(updateQuery, newPhoneNumber, numberId, contactId)
	}

	var retrievedId uuid.UUID
	err := row.Scan(&retrievedId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundErr(errs.NumberNotFoundErr)
		}
		logger.Error("Error while updating record in numbers table: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	return &newNumber, nil
}

func (r ContactRepositoryDb) UpdateContactEmail(newEmail Email) (*Email, *errs.AppError) {
	emailId := newEmail.Id
	contactId := newEmail.ContactId
	newAddress := newEmail.Address

	if appErr := r.CheckContactExistenceById(contactId); appErr != nil {
		return nil, appErr
	}

	updateQuery := `UPDATE emails SET address = $1 WHERE id = $2 AND contact_id = $3 RETURNING id`
	row := r.client.QueryRow(updateQuery, newAddress, emailId, contactId)
	var retrievedId uuid.UUID
	err := row.Scan(&retrievedId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundErr("email with this id not found")
		}
		logger.Error("Error while updating record in emails table: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	return &newEmail, nil
}

func (r ContactRepositoryDb) UpdateContact(newContact Contact) (*Contact, *errs.AppError) {
	contactId := newContact.Id
	newFirstname := newContact.FirstName
	newLastname := newContact.LastName

	if appErr := r.CheckContactExistenceById(contactId); appErr != nil {
		return nil, appErr
	}

	var row *sql.Row
	if newFirstname != "" && newLastname != "" {
		updateQuery := `UPDATE contacts SET first_name = $1, last_name = $2 WHERE id = $3 RETURNING id`
		row = r.client.QueryRow(updateQuery, newFirstname, newLastname, contactId)
	} else if newFirstname != "" && newLastname == "" {
		updateQuery := `UPDATE contacts SET first_name = $1 WHERE id = $2 RETURNING id`
		row = r.client.QueryRow(updateQuery, newFirstname, contactId)
	} else if newFirstname == "" && newLastname != "" {
		updateQuery := `UPDATE contacts SET last_name = $1 WHERE id = $2 RETURNING id`
		row = r.client.QueryRow(updateQuery, newLastname, contactId)
	}

	var retrievedId uuid.UUID
	err := row.Scan(&retrievedId)
	if err != nil {
		logger.Error("Error while updating record in contact tables: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	return &newContact, nil
}
