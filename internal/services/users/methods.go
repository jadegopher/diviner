package users

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/model"
)

func (u *users) AddNewUser(update tgbotapi.Update) (*model.User, error) {
	insert := &model.User{
		Id:           uint64(update.Message.Chat.ID),
		Username:     update.Message.Chat.UserName,
		FirstName:    update.Message.Chat.FirstName,
		LastName:     update.Message.Chat.LastName,
		Phone:        "",
		SaveLocation: false,
	}
	res, err := u.db.User().Select(insert.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if res, err = u.db.User().Insert(*insert); err != nil {
				return nil, err
			}
		}
	}
	return res, nil
}
