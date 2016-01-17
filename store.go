package ppc

import (
	"database/sql"

	"github.com/bentranter/password"
	// "github.com/lib/pq"
)

var defaultStore = newDB()

// PqStore is a reference to the password store.
type PqStore struct {
	DB *sql.DB
}

// Store stores the given id and secret in Postgres. It will hash the secret
// using bcrypt before storing it.
func (p *PqStore) Store(id string, secret string) (string, error) {
	hashedSecret, err := password.Hash(secret)
	if err != nil {
		return "", err
	}

	// @TODO
	p.DB.Query("", hashedSecret)
	return "", nil
}

// Retrieve retrieves from Postgres the hashed secret given an id and secret.
func (p *PqStore) Retrieve(id string, secret string) (string, error) {
	p.DB.Query("")

	return "", nil
}

// newDB returns a default DB
func newDB() *PqStore {
	db, err := sql.Open("postgres", "user=ppc dbname=ppc sslmode=verify-full")
	if err != nil {
		panic(err)
	}

	return &PqStore{
		DB: db,
	}
}
