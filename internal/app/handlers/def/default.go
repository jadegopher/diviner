package def

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"telegram-pug/internal/app/handlers/def/messages"
	"telegram-pug/internal/app/keyboards"
	"telegram-pug/internal/services/users"
	"telegram-pug/repo"
	"telegram-pug/usecases"
	"time"
)

type def struct {
	userService usecases.IUserService
}

func New(dbConn *gorm.DB) (repo.IHandler, error) {
	db, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &def{userService: db}, nil
}

func (d *def) Condition(update tgbotapi.Update) bool {
	if keyboards.MenuKeyboard.IsKeyboardButton(update.Message.Text) {
		return false
	}
	if keyboards.LanguageKeyboard.IsKeyboardButton(update.Message.Text) {
		return false
	}
	return true
}

func (d *def) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	user, err := d.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}
	user.LastTimeOnline = time.Now().UTC().Unix()
	user, err = d.userService.UpdateUserInfo(user)
	if err != nil {
		return nil, err
	}
	log.Println(update.Message.Chat, update.Message.Text, update.ChosenInlineResult)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages.DefaultResponse.CreateResponse(user.Language))

	if update.Message.Text[len(update.Message.Text)-1] == '?' {
		r := rand.Intn(2)
		if r == 0 {
			msg.Text = messages.NoResponse.CreateResponse(user.Language)
		}
		msg.Text = messages.YesResponse.CreateResponse(user.Language)
	}

	msg.ReplyMarkup = keyboards.MenuKeyboard.Keyboard(user.Language)
	return &msg, nil
}
