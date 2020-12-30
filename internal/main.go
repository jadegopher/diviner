package main

import (
	"github.com/Syfaro/telegram-bot-api"
	config "telegram-pug/config"
	"telegram-pug/internal/database/postgres/pq"
	"telegram-pug/internal/handler"
	"telegram-pug/internal/keyboard"
)

func main() {
	c := config.New("config/dev.env")
	if err := c.Read(); err != nil {
		panic(err)
	}

	dbConn, err := pq.NewConnection(c.PostgresLogin(), c.PostgresPassword(),
		c.PostgresHost(), c.PostgresPort(), c.PostgresDbName())
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(c.Token())
	if err != nil {
		panic(err)
	}

	h := handler.New(dbConn, bot, keyboard.New())

	if err := h.HandleUpdates(); err != nil {
		panic(err)
	}
}
