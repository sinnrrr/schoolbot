package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sinnrrr/schoolbot/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	bot, err := tb.NewBot(config.BotSettings)
	if err != nil {
		panic(err)
	}

	println(bot)

	bot.Handle("/hello", func(msg *tb.Message) {
		bot.Send(msg.Sender, "Hello world")
	})

	bot.Start()
}
