package database

import (
	"gorm.io/gorm"
	"telegram-pug/internal/database/postgres/gorm/coordinates"
	"telegram-pug/internal/database/postgres/gorm/user"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type client struct {
	user        repo.IUser
	coordinates repo.ICoordinates
}

func New(db *gorm.DB) (usecases.IHandlerDB, error) {
	ret := &client{
		user:        user.New(db),
		coordinates: coordinates.New(db),
	}
	if err := ret.User().CreateTableIfNotExists(); err != nil {
		return nil, err
	}
	if err := ret.Coordinates().CreateTableIfNotExists(); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *client) User() repo.IUser {
	return c.user
}

func (c *client) Coordinates() repo.ICoordinates {
	return c.coordinates
}
