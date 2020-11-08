package main

import (
	"flag"
	"github.com/chai2010/gettext-go"
	"github.com/sinnrrr/schoolbot/config"
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
)

var (
	bot *tb.Bot
	l = gettext.New("general", "locale").SetLanguage(os.Getenv("DEFAULT_LANGUAGE"))
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}

	isCron := flag.Bool("cron", false, "is the cron job?")
	flag.Parse()

	if *isCron {
		cronAlerts()
	} else {
		err = InitTunnel()
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

		err = syncSupportedLanguage()
		if err != nil {
			panic(err)
		}

		registerInlineKeyboard()
		registerKeyboard()

		handleOnTextEvent()
		handleOnAddedEvent()

		handleStartCommand()

		log.Println("Websocket has been set up on", os.Getenv("PUBLIC_URL"))
		log.Println("Bot has been started on port", os.Getenv("PORT"))

		bot.Start()
	}
}
