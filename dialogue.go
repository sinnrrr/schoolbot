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
	homework = make(map[string]interface{})
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
					m.Sender,
					"Sorry, but I don't understand you. Їм сало...",
				),
			)
		case SubjectRequest:
			homework["subject"] = m.Text
			err := db.SetDialogueState(m.Sender.ID, HomeworkRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Sender,
					"Enter task now",
				),
			)
		case HomeworkRequest:
			homework["task"] = m.Text
			err := db.SetDialogueState(m.Sender.ID, NoRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Sender,
					"Homework from subject was successfuly created",
					keyboard,
				),
			)
		}
	})
}
