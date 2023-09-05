package core

// Number core object definition
type Number struct {
	Id          string `db:"id"`
	ContactId   string `db:"contact_id"`
	PhoneNumber string `db:"number"`
	Label       string `db:"label"`
}
