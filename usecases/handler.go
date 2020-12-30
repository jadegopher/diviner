package usecases

import "telegram-pug/repo"

type IHandlerDB interface {
	User() repo.IUser
	Coordinates() repo.ICoordinates
}
