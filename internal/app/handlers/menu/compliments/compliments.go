package compliments

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/constants"
	"telegram-pug/internal/app/handlers/menu/compliments/messages"
	"telegram-pug/internal/database/errs"
	complimentsDb "telegram-pug/internal/database/postgres/gorm/compliments"
	"telegram-pug/internal/services/users"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type compliments struct {
	userService       usecases.IUserService
	complimentService model.ICompliments
}

func New(dbConn *gorm.DB, path string) (repo.IHandler, error) {
	userDb, err := users.New(dbConn)
	if err != nil {
		return nil, err
	}
	ret := &compliments{
		complimentService: complimentsDb.New(dbConn, path),
		userService:       userDb,
	}
	if err = ret.complimentService.CreateTableIfNotExists(); err != nil {
		return nil, err
	}
	return ret, nil
}

func (w *compliments) Handle(update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	user, err := w.userService.UserInfo(update)
	if err != nil {
		return nil, err
	}

	compliment, err := w.complimentService.SelectRandom(model.Neutral.String())
	if err != nil && err != errs.RecordNotFound {
		return nil, err
	} else if err != nil && err == errs.RecordNotFound {
		switch user.Language {
		case constants.Eng:
			msg.Text = messages.ComplimentErr.English()
		case constants.Ru:
			msg.Text = messages.ComplimentErr.Russian()
		default:
			msg.Text = messages.ComplimentErr.English()
		}
		return &msg, nil
	} else if err != nil {
		return nil, err
	}

	switch user.Language {
	case constants.Eng:
		msg.Text = compliment.English
	case constants.Ru:
		msg.Text = compliment.Russian
	default:
		msg.Text = compliment.English
	}

	return &msg, nil
}
