package config

import (
	"github.com/afdalwahyu/gonnel"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

var (
	TunnelOptions = gonnel.Options{
		BinaryPath: os.Getenv("NGROK_PATH"),
	}

	Tunnel = gonnel.Tunnel{
		Proto:        gonnel.HTTP,
		Name:         "schoolbot",
		LocalAddress: ":" + os.Getenv("PORT"),
	}

	BotWebhook = &tb.Webhook{
		Listen: ":" + os.Getenv("PORT"),
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: os.Getenv("PUBLIC_URL"),
		},
	}

	BotSettings = tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 15 * time.Second},
	}

	DBLogLevel neo4j.LogLevel = neo4j.INFO

	DB = func() func(config *neo4j.Config) {
		return func(config *neo4j.Config) {
			config.Encrypted = false
			config.Log = neo4j.ConsoleLogger(DBLogLevel)
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