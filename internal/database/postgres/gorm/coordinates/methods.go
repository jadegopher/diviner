package coordinates

import (
	"errors"
	"telegram-pug/model"
)

func (c *coordinates) CreateTableIfNotExists() error {
	if err := c.db.AutoMigrate(model.Coordinates{}); err != nil {
		return err
	}
	return nil
}

func (c *coordinates) Select(userId uint) (model.Coordinates, error) {
	return model.Coordinates{}, errors.New("method not implemented")
}

func (c *coordinates) SelectMany(userId, limit, offset uint) ([]model.Coordinates, error) {
	return nil, errors.New("method not implemented")
}

func (c *coordinates) Insert(coordinates model.Coordinates) (model.Coordinates, error) {
	return model.Coordinates{}, errors.New("method not implemented")
}

func (c *coordinates) Delete(userId uint) error {
	return errors.New("method not implemented")
}
