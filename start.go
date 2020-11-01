package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func handleStartCommand() {
	bot.Handle("/start", func(m *tb.Message) {
		if m.Private() {
			if m.Payload != "" {
				classID, err := strconv.ParseInt(m.Payload, 10, 64)
				if err != nil {
					panic(err)
				}

				student, err := db.CreateStudent(m.Sender, classID)
				if err != nil {
					panic(err)
				}

				if student == nil {
					handleSendError(
						bot.Send(
							m.Sender,
							"You have already accepted the invite from this group",
							keyboard,
						),
					)
				} else {
					handleSendError(
						bot.Send(
							m.Chat,
							"Hello, how can I help?",
							keyboard,
						),
					)
				}
			} else {
				handleSendError(
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
			handleSendError(
				bot.Send(
					m.Chat,
					"Hello, how can I help in your group?",
					keyboard,
				),
			)
		}
	})
}

func handleOnAddedEvent() {
	bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		class, err := db.CreateClass(m.Chat.ID, m.Chat.Title)
		if err != nil {
			panic(err)
		}

		if class == nil {
			handleSendError(
				bot.Send(
					m.Chat,
					"Your group have already records in our database",
				),
			)
		} else {
			handleSendError(
				bot.Send(
					m.Chat,
					"Invite to your personal chat",
					&tb.ReplyMarkup{
						InlineKeyboard: generatePersonalInviteKeys(m.Chat.ID),
					},
				),
			)
		}
	})
}

func handleSendError(m *tb.Message, err error) {
	if err != nil {
		bot.Send(m.Chat, "Something went wrong with bot. Please, try again later.")
		panic(err)
	}
}
