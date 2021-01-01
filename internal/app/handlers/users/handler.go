package users

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/model"
)

func (u *users) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	if update.Message.Chat.Type == "private" {
		if err := u.AddNewUser(update); err != nil {
			return nil, err
		}
	}
	fmt.Println(update.Message.Chat)
	return &msg, nil
}

func (u *users) AddNewUser(update tgbotapi.Update) error {
	insert := model.User{
		Id:           uint64(update.Message.Chat.ID),
		Username:     update.Message.Chat.UserName,
		FirstName:    update.Message.Chat.FirstName,
		LastName:     update.Message.Chat.LastName,
		Phone:        "",
		SaveLocation: false,
	}
	if _, err := u.db.User().Select(insert.Id); err != nil {
		if err == gorm.ErrRecordNotFound {
			if _, err := u.db.User().Insert(insert); err != nil {
				return err
			}
		}
	}
	return nil
}
