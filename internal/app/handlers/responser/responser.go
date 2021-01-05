package responser

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"telegram-pug/internal/app/handlers/responser/messages"
	"telegram-pug/internal/app/keyboards"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type responser struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.ICondition, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &responser{userService: db}, nil
}

func (r *responser) Condition(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	if keyboards.MenuKeyboard.IsKeyboardButton(update.Message.Text) {
		return nil, nil
	}
	if keyboards.LanguageKeyboard.IsKeyboardButton(update.Message.Text) {
		return nil, nil
	}
	if len(update.Message.Text) == 0 {
		return nil, nil
	}
	return r.Handle(update)
}

func (r *responser) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	user, err := r.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	log.Println(update.Message.Chat, update.Message.Text, update.ChosenInlineResult)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.DefaultResponse.CreateResponse(user.Language))

	if update.Message.Text[len(update.Message.Text)-1] == '?' && update.Message.Text != "?" {
		r := rand.Intn(2)
		if r == 0 {
			msg.Text = messages.NoResponse.CreateResponse(user.Language)
		} else {
			msg.Text = messages.YesResponse.CreateResponse(user.Language)
		}
	}

	msg.ReplyMarkup = keyboards.MenuKeyboard.Keyboard(user.Language)
	return &msg, nil
}
