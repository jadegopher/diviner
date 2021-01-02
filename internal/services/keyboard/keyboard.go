package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-pug/constants"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type keyboard struct {
	buttons [][]repo.IButton
}

func New(buttons ...[]repo.IButton) usecases.IKeyboard {
	var ret [][]repo.IButton
	ret = append(ret, buttons...)
	return &keyboard{buttons: ret}
}

func (k *keyboard) IsKeyboardButton(button string) bool {
	for _, row := range k.buttons {
		for _, btn := range row {
			if btn.English() == button || btn.Russian() == button {
				return true
			}
		}
	}
	return false
}

func (k *keyboard) KeyboardButton(button string) (repo.IButton, bool) {
	for _, row := range k.buttons {
		for _, btn := range row {
			if btn.English() == button || btn.Russian() == button {
				return btn, true
			}
		}
	}
	return nil, false
}

func (k *keyboard) Keyboard(language string) tgbotapi.ReplyKeyboardMarkup {
	var tgKeyboard [][]tgbotapi.KeyboardButton
	for _, row := range k.buttons {
		var tmp []tgbotapi.KeyboardButton
		for _, btn := range row {
			text := ""
			switch language {
			case constants.Eng:
				text = btn.English()
			case constants.Ru:
				text = btn.Russian()
			default:
				text = btn.English()
			}
			if btn.Type() == model.Default {
				tmp = append(tmp, tgbotapi.NewKeyboardButton(text))
			} else if btn.Type() == model.Location {
				tmp = append(tmp, tgbotapi.NewKeyboardButtonLocation(text))
			}
		}
		tgKeyboard = append(tgKeyboard, tmp)
	}
	return tgbotapi.NewReplyKeyboard(tgKeyboard...)
}
