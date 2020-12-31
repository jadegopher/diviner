package handler

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
	"log"
	"telegram-pug/internal/handler/database"
	"telegram-pug/internal/handler/weather"
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

	for update := range updates {
		if update.Message == nil {
			continue
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			fmt.Println(update.Message.Chat)
			fmt.Println(update.Message.Date)
			if update.Message.Location != nil {
				msg, err = weather.Handle(h.weatherToken, update)
				if err != nil {
					log.Println(err)
				}
			}
			switch update.Message.Text {
			case "/start":
				msg.Text = "Henlo stranger... I'm the diviner pug. I forgot everything except my name... " +
					"I can't recall what happened with me but I'm trying... I lost skills of the greatest diviner.... " +
					"Oh wow, there is a scroll of the acsients with some magic stuff. Hmmmm, may be I could try to do " +
					"some street magic? But I need your help.Tap the button when you will be ready"
				msg.ReplyMarkup = h.keyboard
			}
			h.bot.Send(msg)
		}
	}
	return nil
}
