package main

import (
	"fmt"
	"github.com/sinnrrr/schoolbot/handlers"
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
		state, err := handlers.GetDialogueState(m.Sender.ID)
		if err != nil {
			panic(err)
		}

		fmt.Println("Got state")
		fmt.Println(state)

		switch state {
		case NoRequest:
			handleSendError(
				bot.Send(
					m.Sender,
					"Sorry, but I don't understand you. Їм сало...",
				),
			)
		case SubjectRequest:
			homework["subject"] = m.Text
			err := handlers.SetDialogueState(m.Sender.ID, HomeworkRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Edit(
					m,
					"Enter task",
				),
			)
		case HomeworkRequest:
			homework["task"] = m.Text
			err := handlers.SetDialogueState(m.Sender.ID, NoRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Edit(
					m,
					"Homework from subject " +
						homework["subject"].(string) +
						" was successfully created",
				),
			)
		}
	})
}
