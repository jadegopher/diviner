package def

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/internal/app/handlers/def/messages"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
	"time"
)

type def struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.IHandler, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &def{userService: db}, nil
}

func (d *def) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	user, err := d.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	user.LastTimeOnline = time.Now().UTC().Unix()
	user, err = d.userService.UpdateUserInfo(user)
	if err != nil {
		return nil, err
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.Response(user.Language))
	return &msg, nil
}
