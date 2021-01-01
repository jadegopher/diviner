package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"telegram-pug/internal/app/handlers/def"
	"telegram-pug/internal/app/handlers/start"
	"telegram-pug/internal/app/handlers/weather"
	"telegram-pug/repo"
)

type handler struct {
	bot      *tgbotapi.BotAPI
	handlers []repo.IHandler
}

func New(dbConn *gorm.DB, bot *tgbotapi.BotAPI, keyboard tgbotapi.ReplyKeyboardMarkup,
	weatherToken string) (*handler, error) {

	defaultHandler, err := def.New(dbConn)
	if err != nil {
		return nil, err
	}

	initHandler, err := start.New(dbConn, keyboard)
	if err != nil {
		return nil, err
	}
	weatherHandler := weather.New(weatherToken)

	handlers := []repo.IHandler{defaultHandler, initHandler, weatherHandler}

	return &handler{
		bot:      bot,
		handlers: handlers,
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

	for update := range updates {
		if update.Message == nil {
			continue
		} else {
			msgs := make([]tgbotapi.MessageConfig, 0, 10)
			for _, handler := range h.handlers {
				msg, err := handler.Handle(update)
				if err == nil && msg != nil {
					msgs = append(msgs, *msg)
				} else {
					log.Println(err)
				}
			}

			for _, message := range msgs {
				if _, err := h.bot.Send(message); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
