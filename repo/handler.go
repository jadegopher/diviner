package repo

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type IHandler interface {
	Handle(update tgbotapi.Update) (tgbotapi.MessageConfig, error)
}
