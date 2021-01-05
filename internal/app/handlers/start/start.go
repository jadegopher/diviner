package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/internal/app/handlers/start/messages"
	"telegram-pug/internal/app/keyboards"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type start struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.ICondition, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &start{userService: db}, nil
}

func (s *start) Condition(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	if update.Message.Text != "/start" {
		return nil, nil
	}
	return s.Handle(update)
}

func (s *start) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	user, err := s.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.ChooseLanguage.English(user.FirstName))
	msg.ReplyMarkup = keyboards.LanguageKeyboard.Keyboard(user.Language)
	return &msg, nil
}
