package user

import (
	"database/sql"
	"telegram-pug/repo"
)

type user struct {
	db *sql.DB
}

func New(db *sql.DB) repo.IUser {
	return &user{db: db}
}
