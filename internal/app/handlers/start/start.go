package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/internal/app/handlers/start/messages"
	"telegram-pug/internal/services/users"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type start struct {
	userService usecases.IUserService
	keyboard    tgbotapi.ReplyKeyboardMarkup
}

func New(dbConn *gorm.DB, keyboard tgbotapi.ReplyKeyboardMarkup) (repo.IHandler, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &start{userService: db, keyboard: keyboard}, nil
}

func (s *start) Condition(update tgbotapi.Update) bool {
	if update.Message.Text != "/start" {
		return false
	}
	return true
}

func (s *start) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {

	user := &model.User{}
	var err error
	user, err = s.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.Text = messages.InitConversation.English(user.Username)
	msg.ReplyMarkup = s.keyboard
	return &msg, nil
}
