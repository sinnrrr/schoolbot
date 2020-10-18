package config

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

var (
	BotWebhook = &tb.Webhook{
		Listen: ":" + os.Getenv("PORT"),
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: os.Getenv("PUBLIC_URL"),
		},
	}

	BotSettings = tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: BotWebhook,
	}

	DB = func(level neo4j.LogLevel) func(config *neo4j.Config) {
		return func(config *neo4j.Config) {
			config.Encrypted = false
			config.Log = neo4j.ConsoleLogger(level)
		}
	}

	Session = neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	}

	URI = "bolt://" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	Auth = neo4j.BasicAuth(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		"",
	)
)
