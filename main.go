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

	bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		bot.Send(
			m.Chat,
			"Invite",
			&tb.ReplyMarkup{InlineKeyboard: personalInviteKeys},
		)
	})

	bot.Handle("/start", func(m *tb.Message) {
		if m.Private() {
			if m.Payload != "" {
				RegisterMenu(bot)
				bot.Send(m.Sender, "Hello", menu)
			} else {
				bot.Send(m.Sender, "Add me to group", &tb.ReplyMarkup{InlineKeyboard: groupInviteKeys})
			}
		}
	})

	bot.Start()
}
