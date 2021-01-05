package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/internal/app/handlers/menu/compliments"
	"telegram-pug/internal/app/handlers/menu/weather"
	"telegram-pug/internal/app/keyboards"
	"telegram-pug/repo"
)

type menu struct {
	weather     repo.IHandler
	compliments repo.IHandler
}

func New(dbConn *gorm.DB, token string, complimentsPath string) (repo.ICondition, error) {
	weatherHandler, err := weather.New(dbConn, token)
	if err != nil {
		return nil, err
	}
	complimentsHandler, err := compliments.New(dbConn, complimentsPath)
	if err != nil {
		return nil, err
	}
	return &menu{
		weather:     weatherHandler,
		compliments: complimentsHandler,
	}, nil
}

func (m *menu) Condition(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {

	if btn, in := keyboards.MenuKeyboard.KeyboardButton(update.Message.Text); in {
		switch btn.English() {
		case keyboards.LocationBtnEn:
			if update.Message.Location != nil {
				return m.weather.Handle(update)
			}
		case keyboards.ComplimentButtonEn:
			return m.compliments.Handle(update)
		}
	}
	return nil, nil
}
