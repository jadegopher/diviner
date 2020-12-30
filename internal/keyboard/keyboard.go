package keyboard

import "github.com/Syfaro/telegram-bot-api"

type keyboard struct {
}

func New() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation("🌎 Send location"),
	))
}
