package handler

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"telegram-pug/internal/handler/database"
	"telegram-pug/usecases"
)

type handler struct {
	db       usecases.IHandlerDB
	bot      *tgbotapi.BotAPI
	keyboard tgbotapi.ReplyKeyboardMarkup
}

func New(db *sql.DB, bot *tgbotapi.BotAPI, keyboard tgbotapi.ReplyKeyboardMarkup) *handler {
	return &handler{
		db:       database.New(db),
		bot:      bot,
		keyboard: keyboard,
	}
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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			fmt.Println(update.Message)
			if update.Message.Location != nil {
				fmt.Println(update.Message.Location)
			}
			switch update.Message.Text {
			case "/start":
				msg.ReplyMarkup = h.keyboard
			}
			msg.Text = "pong"
			if _, err := h.bot.Send(msg); err != nil {
				return err
			}
		}
	}
	return nil
}
