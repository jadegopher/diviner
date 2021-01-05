package repo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type ICondition interface {
	Condition(update tgbotapi.Update) (*tgbotapi.MessageConfig, error)
}

type IHandler interface {
	Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error)
}
