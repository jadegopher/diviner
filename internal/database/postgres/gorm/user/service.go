package user

import (
	"gorm.io/gorm"
	"telegram-pug/model"
)

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) model.IUser {
	return &user{db: db}
}
