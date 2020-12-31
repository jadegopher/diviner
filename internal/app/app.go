package app

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"telegram-pug/internal/app/database"
	"telegram-pug/internal/app/handlers/start"
	"telegram-pug/internal/app/handlers/weather"
	"telegram-pug/usecases"
)

type handler struct {
	db           usecases.IHandlerDB
	bot          *tgbotapi.BotAPI
	keyboard     tgbotapi.ReplyKeyboardMarkup
	weatherToken string
}

func New(dbConn *gorm.DB, bot *tgbotapi.BotAPI, keyboard tgbotapi.ReplyKeyboardMarkup,
	weatherToken string) (*handler, error) {
	db, err := database.New(dbConn)
	if err != nil {
		return nil, err
	}

	return &handler{
		db:           db,
		bot:          bot,
		keyboard:     keyboard,
		weatherToken: weatherToken,
	}, nil
}

func (h *handler) HandleUpdates() error {
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := h.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	initHandler := start.New(h.keyboard)
	weatherHandler := weather.New(h.weatherToken)

	for update := range updates {
		if update.Message == nil {
			continue
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			fmt.Println(update.Message.Chat)
			fmt.Println(update.Message.Date)

			msg, err = initHandler.Handle(update)
			if err != nil {
				log.Println(err)
			}

			msg, err = weatherHandler.Handle(update)
			if err != nil {
				log.Println(err)
			}

			h.bot.Send(msg)
		}
	}
	return nil
}
