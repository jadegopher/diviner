package languages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func New() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(EngButton),
		tgbotapi.NewKeyboardButton(RuButton),
	))
}
