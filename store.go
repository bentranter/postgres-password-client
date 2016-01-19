package ppc

import (
	"database/sql"

	"github.com/bentranter/password"
	// Use the Postgres DB driver
	_ "github.com/lib/pq"
)

var defaultStore = newDB()

// PgStore is a reference to the password store.
type PgStore struct {
	DB *sql.DB
}

// Store stores the given id and secret in Postgres. It will hash the secret
// using bcrypt before storing it.
func (p *PgStore) Store(id string, secret string) (string, error) {
	hashedSecret, err := password.Hash(secret)
	if err != nil {
		return "", err
	}

	// @TODO
	p.DB.Query("", hashedSecret)
	return "", nil
}

// Retrieve retrieves from Postgres the hashed secret given an id and secret.
func (p *PgStore) Retrieve(id string, secret string) (string, error) {
	p.DB.Query("")

	return "", nil
}

// newDB returns a default DB
func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=ppc dbname=ppc sslmode=verify-full")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}
