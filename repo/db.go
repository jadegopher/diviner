package repo

import (
	"telegram-pug/model"
)

type IDb interface {
	User() model.IUser
	Coordinates() model.ICoordinates
}
