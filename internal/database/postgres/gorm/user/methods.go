package user

import (
	"errors"
	"telegram-pug/model"
)

func (u *user) CreateTableIfNotExists() error {
	if err := u.db.AutoMigrate(model.User{}); err != nil {
		return err
	}
	return nil
}

func (u *user) Select(id uint) (model.User, error) {
	return model.User{}, errors.New("method not implemented")
}

func (u *user) Insert(user model.User) (model.User, error) {
	return model.User{}, errors.New("method not implemented")
}

func (u *user) Update(user model.User) (model.User, error) {
	return model.User{}, errors.New("method not implemented")
}

func (u *user) Delete(id uint) error {
	return errors.New("method not implemented")
}
