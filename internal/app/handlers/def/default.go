package def

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
	"time"
)

type def struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.ICondition, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &def{userService: db}, nil
}

func (d *def) Condition(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	return d.Handle(update)
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
	log.Println(update.Message.Chat, update.Message.Text, update.ChosenInlineResult)

	return nil, nil
}
