package keyboards

import (
	"telegram-pug/internal/services/button"
	"telegram-pug/internal/services/keyboard"
	"telegram-pug/model"
	"telegram-pug/repo"
)

const LocationBtnEn = "🌎 Send location"
const LocationBtnRu = "🌎 Отправить геолокацию"
const ComplimentButtonEn = "💚 Get a compliment"
const ComplimentButtonRu = "💚 Получить комплимент"

var MenuKeyboard = keyboard.New(
	[]repo.IButton{
		button.New(LocationBtnEn, LocationBtnRu, model.Location),
	},
	[]repo.IButton{
		button.New(ComplimentButtonEn, ComplimentButtonRu, model.Default),
	},
)
