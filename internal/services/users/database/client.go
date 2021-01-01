package database

import (
	"gorm.io/gorm"
	"telegram-pug/internal/database/postgres/gorm/coordinates"
	"telegram-pug/internal/database/postgres/gorm/user"
	"telegram-pug/model"
	"telegram-pug/repo"
)

type client struct {
	user        model.IUser
	coordinates model.ICoordinates
}

func New(db *gorm.DB) (repo.IDb, error) {
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

func (c *client) User() model.IUser {
	return c.user
}

func (c *client) Coordinates() model.ICoordinates {
	return c.coordinates
}
