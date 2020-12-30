package coordinates

import (
	"database/sql"
	"telegram-pug/repo"
)

type coordinates struct {
	db *sql.DB
}

func New(db *sql.DB) repo.ICoordinates {
	return &coordinates{db: db}
}
