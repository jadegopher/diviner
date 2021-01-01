package repo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type IHandler interface {
	Condition(update tgbotapi.Update) bool
	Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error)
}
