package usecases

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-pug/repo"
)

type IKeyboard interface {
	IsKeyboardButton(button string) bool
	KeyboardButton(button string) (repo.IButton, bool)
	Keyboard(language string) tgbotapi.ReplyKeyboardMarkup
}
