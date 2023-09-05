package core

// Email core object definition
type Email struct {
	Id        string `db:"id"`
	ContactId string `db:"contact_id"`
	Address   string `db:"address"`
}
