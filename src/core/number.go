package core

// Number core object definition
type Number struct {
	ID        string `db:"id"`
	ContactId string `db:"contact_id"`
	Number    string `db:"number"`
	Label     string `db:"label"`
}

// NumberRepository Secondary port
type NumberRepository interface {
}
