package config

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

var (
	BotSettings = tb.Settings{
		URL: "https://" + os.Getenv("HOST"),
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
)
