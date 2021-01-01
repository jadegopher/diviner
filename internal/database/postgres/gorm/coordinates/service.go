package coordinates

import (
	"gorm.io/gorm"
	"telegram-pug/model"
)

type coordinates struct {
	db *gorm.DB
}

func New(db *gorm.DB) model.ICoordinates {
	return &coordinates{db: db}
}
