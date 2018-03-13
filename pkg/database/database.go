package database

import (
	"database/sql"

	"github.com/knq/dburl"
)

type IDatabaseConnection interface {
	Initialise(map[string]string) (*sql.DB, error)
}
type DatabaseWrapper struct {
	DB *sql.DB
}

func (dw *DatabaseWrapper) Initialise(params map[string]string) (*sql.DB, error) {
	// need to construct this URL from the params map passed
	u, err := dburl.Open("postgresql://local:local@localhost/local?sslmode=disable")
	return u, err
}
