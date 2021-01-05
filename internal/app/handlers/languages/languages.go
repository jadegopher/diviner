package languages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/constants"
	"telegram-pug/internal/app/handlers/languages/messages"
	"telegram-pug/internal/app/keyboards"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
	"time"
)

type lang struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.ICondition, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &lang{userService: db}, nil
}

func (l *lang) Condition(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	if !keyboards.LanguageKeyboard.IsKeyboardButton(update.Message.Text) {
		return nil, nil
	}
	return l.Handle(update)
}

func (l *lang) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	user, err := l.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	user.LastTimeOnline = time.Now().UTC().Unix()
	if button, in := keyboards.LanguageKeyboard.KeyboardButton(update.Message.Text); in {
		switch button.English() {
		case keyboards.EngButton:
			user.Language = constants.Eng
		case keyboards.RuButton:
			user.Language = constants.Ru
		default:
			user.Language = constants.Eng
		}
	}
	user, err = l.userService.UpdateUserInfo(user)
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.InitConversation.CreateResponse(user.Language,
		user.FirstName))

	msg.ReplyMarkup = keyboards.MenuKeyboard.Keyboard(user.Language)
	return &msg, nil
}
