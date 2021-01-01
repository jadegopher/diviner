package user

import (
	"errors"
	"fmt"
	"reflect"
	"telegram-pug/model"
)

func (u *user) CreateTableIfNotExists() error {
	if _, err := u.db.Exec(fmt.Sprintf(`create table if not exists %s`,
		reflect.TypeOf(model.User{}).Name())); err != nil {
		return err
	}
	return nil
}

func (u *user) Select(id uint64) (*model.User, error) {
	return nil, errors.New("method not implemented")
}

func (u *user) Insert(user model.User) (*model.User, error) {
	return nil, errors.New("method not implemented")
}

func (u *user) Update(user model.User) (*model.User, error) {
	return nil, errors.New("method not implemented")
}

func (u *user) Delete(id uint64) error {
	return errors.New("method not implemented")
}
