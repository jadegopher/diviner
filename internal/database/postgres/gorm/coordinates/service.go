package coordinates

import (
	"gorm.io/gorm"
	"telegram-pug/repo"
)

type coordinates struct {
	db *gorm.DB
}

func New(db *gorm.DB) repo.ICoordinates {
	return &coordinates{db: db}
}
