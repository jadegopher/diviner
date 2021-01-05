package compliments

import (
	"gorm.io/gorm"
	"telegram-pug/model"
)

type compliments struct {
	dataPath string
	db       *gorm.DB
}

func New(db *gorm.DB, path string) model.ICompliments {
	return &compliments{db: db, dataPath: path}
}
