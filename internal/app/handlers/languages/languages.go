package languages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/constants"
	"telegram-pug/internal/app/handlers/languages/messages"
	"telegram-pug/internal/services/keyboards/languages"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
	"time"
)

type lang struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.IHandler, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &lang{userService: db}, nil
}

func (l *lang) Condition(update tgbotapi.Update) bool {
	if update.Message.Text != languages.EngButton && update.Message.Text != languages.RuButton {
		return false
	}
	return true
}

func (l *lang) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	user, err := l.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	user.LastTimeOnline = time.Now().UTC().Unix()
	switch update.Message.Text {
	case languages.EngButton:
		user.Language = constants.Eng
	case languages.RuButton:
		user.Language = constants.Ru
	default:
		user.Language = constants.Eng
	}
	user, err = l.userService.UpdateUserInfo(user)
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.InitConversation.CreateResponse(user.Language,
		user.FirstName))
	return &msg, nil
}
