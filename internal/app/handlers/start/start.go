package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/internal/app/handlers/start/database"
	"telegram-pug/internal/app/handlers/start/messages"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type start struct {
	db       usecases.IHandlerDB
	keyboard tgbotapi.ReplyKeyboardMarkup
}

func New(dbConn *gorm.DB, keyboard tgbotapi.ReplyKeyboardMarkup) (repo.IHandler, error) {
	db, err := database.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &start{db: db, keyboard: keyboard}, nil
}

func (s *start) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	if update.Message.Text == "/start" {
		user := &model.User{}
		var err error
		if update.Message.Chat.Type == "private" {
			user, err = s.db.User().Select(uint64(update.Message.Chat.ID))
			if err != nil {
				return nil, err
			}
		}
		msg.Text = messages.InitConversation.English(user.Username)
		msg.ReplyMarkup = s.keyboard
	}
	return &msg, nil
}
