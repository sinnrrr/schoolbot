package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sinnrrr/schoolbot/config"
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

var bot *tb.Bot

func main() {
	db.Init()

	err := InitTunnel()
	if err != nil {
		panic(err)
	}

	bot, err = tb.NewBot(config.BotSettings)
	if err != nil {
		panic(err)
	}

	err = bot.SetWebhook(config.BotWebhook)
	if err != nil {
		panic(err)
	}

	registerKeyboard()
	registerInlineKeyboard()

	handleStartCommand()
	handleOnAddedEvent()

	println("Websocket has been set up on", os.Getenv("PUBLIC_URL"))
	println("Bot has been started on port", os.Getenv("PORT"))

	bot.Start()
}
