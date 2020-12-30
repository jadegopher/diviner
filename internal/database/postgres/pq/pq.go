package pq

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewConnection(user, password, host, port, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s", dbName, user, password, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
