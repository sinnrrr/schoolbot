package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func handleStartCommand() {
	l.SetDomain("dialogue")

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
							l.Gettext("You have already accepted the invite from this group :p"),
							keyboard,
						),
					)
				} else {
					handleSendError(
						bot.Send(
							m.Chat,
							l.Gettext("Hello, how can I help my good old friend? :)"),
							keyboard,
						),
					)
				}
			} else {
				handleSendError(
					bot.Send(
						m.Chat,
						l.Gettext("To get things started, please add me to your class group :]"),
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
					l.Gettext("Hello, how can I help in your group?"),
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
					l.Gettext("Your group have already records in our database"),
				),
			)
		} else {
			handleSendError(
				bot.Send(
					m.Chat,
					l.Gettext("Hey guys! Click this button in order to have access to create and read homeworks and alerts from ths group ;p"),
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
		bot.Send(m.Chat, l.Gettext("Something went wrong with bot. Please, try again later."))
		panic(err)
	}
}
