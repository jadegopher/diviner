package keyboards

import (
	"telegram-pug/internal/services/button"
	"telegram-pug/internal/services/keyboard"
	"telegram-pug/model"
	"telegram-pug/repo"
)

const EngButton = "🇬🇧 English"
const RuButton = "🇷🇺 Русский"

var LanguageKeyboard = keyboard.New([]repo.IButton{
	button.New(EngButton, EngButton, model.Default),
},
	[]repo.IButton{
		button.New(RuButton, RuButton, model.Default),
	},
)
