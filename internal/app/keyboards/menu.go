package keyboards

import (
	"telegram-pug/internal/services/button"
	"telegram-pug/internal/services/keyboard"
	"telegram-pug/model"
	"telegram-pug/repo"
)

const LocationBtnEn = "🌎 Send location"
const LocationBtnRu = "🌎 Отправить геолокацию"

var MenuKeyboard = keyboard.New([]repo.IButton{
	button.New("🌎 Send location", "🌎 Отправить геолокацию", model.Location),
})
