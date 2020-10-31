package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	NoRequest       = 0
	SubjectRequest  = 1
	HomeworkRequest = 2
)

var (
	newHomework = make(map[string]interface{})
)

func handleOnTextEvent() {
	bot.Handle(tb.OnText, func(m *tb.Message) {
		dialogueState, err := db.DialogueState(m.Sender.ID)
		if err != nil {
			panic(err)
		}

		switch dialogueState {
		case NoRequest:
			handleSendError(
				bot.Send(
					m.Chat,
					"Sorry, but I don't understand you. Їм сало...",
				),
			)
		case SubjectRequest:
			newHomework["subject"] = m.Text
			err := db.SetDialogueState(m.Sender.ID, HomeworkRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter task now",
				),
			)
		case HomeworkRequest:
			newHomework["tg_id"] = m.Sender.ID
			newHomework["task"] = m.Text
			err := db.SetDialogueState(m.Sender.ID, NoRequest)
			if err != nil {
				panic(err)
			}

			_, err = db.CreateHomework(newHomework)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Homework from subject was successfully created",
					keyboard,
				),
			)
		}
	})
}
