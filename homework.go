package main

import (
	"github.com/sinnrrr/schoolbot/handlers"
	tb "gopkg.in/tucnak/telebot.v2"
)

func handleHomeworkCommand() {
	bot.Handle("/homework", func(m *tb.Message) {
		homeworks, err := handlers.QueryHomework()
		if err != nil {
			panic(err)
		}

		if homeworks != nil {
			handleSendError(
				bot.Send(
					m.Chat,
					"No homeworks :)",
				),
			)
		}

		handleSendError(
			bot.Send(
				m.Chat,
				"Handled homework command",
			),
		)
	})
}
