package database

import (
	"database/sql"
	"telegram-pug/internal/database/postgres/pq/coordinates"
	"telegram-pug/internal/database/postgres/pq/user"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type client struct {
	user        repo.IUser
	coordinates repo.ICoordinates
}

func New(db *sql.DB) usecases.IHandlerDB {
	return &client{
		user:        user.New(db),
		coordinates: coordinates.New(db),
	}
}

func (c *client) User() repo.IUser {
	return c.user
}

func (c *client) Coordinates() repo.ICoordinates {
	return c.coordinates
}
