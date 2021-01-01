package usecases

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-pug/model"
)

type IUserService interface {
	UserInfo(update tgbotapi.Update) (user *model.User, err error)
	UpdateUserInfo(user *model.User) (*model.User, error)
}
