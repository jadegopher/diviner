package repo

import "telegram-pug/model"

type IUser interface {
	CreateTableIfNotExists() error
	Select(id uint64) (*model.User, error)
	Insert(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(id uint64) error
}
