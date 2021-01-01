package repo

import "telegram-pug/model"

type ICoordinates interface {
	CreateTableIfNotExists() error
	Select(userId uint64) (model.Coordinates, error)
	SelectMany(userId, limit, offset uint64) ([]model.Coordinates, error)
	Insert(coordinates model.Coordinates) (model.Coordinates, error)
	Delete(userId uint64) error
}
