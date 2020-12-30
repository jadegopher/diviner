package repo

import "telegram-pug/model"

type ICoordinates interface {
	CreateTableIfNotExist() error
	Select(userId uint) (model.Coordinates, error)
	SelectMany(userId, limit, offset uint) ([]model.Coordinates, error)
	Insert(coordinates model.Coordinates) (model.Coordinates, error)
	Delete(userId uint) error
}
