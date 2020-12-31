package user

import (
	"gorm.io/gorm"
	"telegram-pug/repo"
)

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) repo.IUser {
	return &user{db: db}
}
