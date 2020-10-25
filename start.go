package main

import tb "gopkg.in/tucnak/telebot.v2"

func handleStartCommand() {
	bot.Handle("/start", func(m *tb.Message) {
		if m.Private() {
			if m.Payload != "" {
				handleBotError(
					bot.Send(
						m.Chat,
						"Hello, how can I help?",
						keyboard,
					),
				)
			} else {
				handleBotError(
					bot.Send(
						m.Chat,
						"To get started, please, add me to group",
						&tb.ReplyMarkup{
							InlineKeyboard: groupInviteKeys,
						},
					),
				)
			}
		} else {
			handleBotError(
				bot.Send(
					m.Chat,
					"Hello, how can I help?",
					keyboard,
				),
			)
		}
	})
}

func handleOnAddedEvent() {
	bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		handleBotError(
			bot.Send(
				m.Chat,
				"Invite to your personal chat",
				&tb.ReplyMarkup{
					InlineKeyboard: personalInviteKeys,
				},
			))
	})
}

func handleBotError(m *tb.Message, err error) {
	if err != nil {
		panic(err)
	}
}
