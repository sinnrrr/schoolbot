package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	NoRequest           = 0
	SubjectRequest      = 1
	TaskRequest         = 2
	TimeRequest         = 3
	AlertContentRequest = 4
)

var (
	newHomework = make(map[string]interface{})
	newAlert    = make(map[string]interface{})
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
			err := db.SetDialogueState(m.Sender.ID, TaskRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter task now",
				),
			)
		case TaskRequest:
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
		case TimeRequest:
			newAlert["time"] = m.Text
			err := db.SetDialogueState(m.Sender.ID, AlertContentRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Now, send me an alert text",
				),
			)
		case AlertContentRequest:
			newAlert["content"] = m.Text
			err := db.SetDialogueState(m.Sender.ID, NoRequest)
			if err != nil {
				panic(err)
			}

			_, err = db.CreateAlert(m.Sender.ID, newAlert)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Alert was successfully created",
					keyboard,
				),
			)
		}
	})
}
