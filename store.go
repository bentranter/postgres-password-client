package ppc

import (
	"database/sql"

	"github.com/bentranter/password"
	// Use the Postgres DB driver
	_ "github.com/lib/pq"
)

var defaultStore = NewDefaultStore()

// PgStore is a reference to the password store.
type PgStore struct {
	DB *sql.DB
}

// Store stores the given id and secret in Postgres. It will hash the secret
// using bcrypt before storing it.
func Store(id string, secret string) (string, error) {
	hashedSecret, err := password.Hash(secret)
	if err != nil {
		return "", err
	}

	var genID string
	err = defaultStore.DB.QueryRow("INSERT INTO users(username, secret) VALUES($1, $2) RETURNING id", id, hashedSecret).Scan(&genID)
	return genID, err
}

// Retrieve retrieves from Postgres the hashed secret given an id and secret.
func Retrieve(id string, secret string) (string, error) {
	var hashedPassword string
	err := defaultStore.DB.QueryRow("SELECT secret FROM users WHERE id = $1", id).Scan(&hashedPassword)
	return hashedPassword, err
}

// NewDefaultStore returns a default DB
func NewDefaultStore() *PgStore {
	db, err := sql.Open("postgres", "user=bentranter dbname=users sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}
