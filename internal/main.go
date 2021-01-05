package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	config "telegram-pug/config"
	"telegram-pug/internal/app"
	"telegram-pug/internal/database/postgres/gorm"
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

	h, err := app.New(dbConn, bot, c.WeatherToken())
	if err != nil {
		panic(err)
	}

	if err = h.HandleUpdates(); err != nil {
		panic(err)
	}
}
