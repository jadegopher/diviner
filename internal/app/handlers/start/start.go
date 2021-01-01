package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-pug/internal/app/handlers/start/messages"
	"telegram-pug/repo"
)

type start struct {
	keyboard tgbotapi.ReplyKeyboardMarkup
}

func New(keyboard tgbotapi.ReplyKeyboardMarkup) repo.IHandler {
	return &start{keyboard: keyboard}
}

func (s *start) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	if update.Message.Text == "/start" {
		msg.Text = messages.InitConversation.English()
		msg.ReplyMarkup = s.keyboard
	}
	return &msg, nil
}
