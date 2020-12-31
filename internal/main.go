package main

import (
	"github.com/Syfaro/telegram-bot-api"
	config "telegram-pug/config"
	"telegram-pug/internal/app"
	"telegram-pug/internal/database/postgres/gorm"
	"telegram-pug/internal/keyboard"
)

func main() {
	c := config.New("config/dev.env")
	if err := c.Read(); err != nil {
		panic(err)
	}

	dbConn, err := gorm.NewConnection(c.PostgresLogin(), c.PostgresPassword(),
		c.PostgresHost(), c.PostgresPort(), c.PostgresDbName())
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(c.TelegramToken())
	if err != nil {
		panic(err)
	}

	h, err := app.New(dbConn, bot, keyboard.New(), c.WeatherToken())
	if err != nil {
		panic(err)
	}

	if err := h.HandleUpdates(); err != nil {
		panic(err)
	}
}
