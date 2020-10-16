package config

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

var (
	BotWebhook = &tb.Webhook{
		Listen: ":" + os.Getenv("PORT"),
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: "https://" + os.Getenv("PUBLIC_URL") + "/",
		},
	}

	BotSettings = tb.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: BotWebhook,
	}
)
