package keyboard

import "github.com/go-telegram-bot-api/telegram-bot-api"

type keyboard struct {
}

func New() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation("🌎 Send location"),
	))
}
