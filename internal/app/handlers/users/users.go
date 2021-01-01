package users

import (
	"gorm.io/gorm"
	"telegram-pug/internal/app/handlers/users/database"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type users struct {
	db usecases.IHandlerDB
}

func New(dbConn *gorm.DB) (repo.IHandler, error) {
	db, err := database.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &users{db: db}, nil
}
