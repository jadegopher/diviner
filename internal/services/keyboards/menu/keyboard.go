package menu

import "github.com/go-telegram-bot-api/telegram-bot-api"

func New() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation(LocationButton),
	))
}
