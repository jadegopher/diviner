package users

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"telegram-pug/internal/services/users/database"
	"telegram-pug/internal/services/users/errs"
	"telegram-pug/model"
	"telegram-pug/repo"
	"telegram-pug/usecases"
)

type users struct {
	db repo.IDb
}

func New(dbConn *gorm.DB) (usecases.IUserService, error) {
	db, err := database.New(dbConn)
	if err != nil {
		return nil, err
	}
	return &users{db: db}, nil
}

func (u *users) UserInfo(update tgbotapi.Update) (user *model.User, err error) {
	if update.Message.Chat.Type != "private" {
		return nil, errs.NotPrivateMsg
	}
	if user, err = u.AddNewUser(update); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *users) UpdateUserInfo(user *model.User) (*model.User, error) {
	if _, err := u.db.User().Select(user.Id); err != nil {
		return nil, err
	}
	ret, err := u.db.User().Update(*user)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
