package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	config "telegram-pug/config"
)

func main() {
	//Создаем бота
	c := config.New("config/dev.env")
	if err := c.Read(); err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(c.Token())
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation("🌎 Send location"),
	))

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
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
				msg.ReplyMarkup = keyboard
			}
			bot.Send(msg)
		}
	}
}
