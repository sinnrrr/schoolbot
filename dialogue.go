package main

import (
	"github.com/sinnrrr/schoolbot/db"
	"github.com/sinnrrr/schoolbot/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

const (
	NoRequest           = 0
	SubjectRequest      = 1
	TaskRequest         = 2
	TimeRequest         = 3
	AlertContentRequest = 4

	MondayRequest    = 5
	TuesdayRequest   = 6
	WednesdayRequest = 7
	ThursdayRequest  = 8
	FridayRequest    = 9
	SaturdayRequest  = 10

	ScheduleRequest = 11
)

var (
	newTimetable = make(map[time.Weekday][]string)
	newHomework  = make(map[string]interface{})
	newAlert     = make(map[string]interface{})

	newSchedule     []string
	createdSchedule map[string]interface{}
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
		case ScheduleRequest:
			newSchedule = utils.ParseSchedule(m.Text)

			err := db.SetDialogueState(m.Sender.ID, MondayRequest)
			if err != nil {
				panic(err)
			}

			createdSchedule, err = db.CreateSchedule(m.Sender.ID, newSchedule)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Sender,
					"Enter lessons for monday",
				),
			)
		case MondayRequest:
			newTimetable[time.Monday] = utils.ParseSubjects(m.Text)

			err := db.SetDialogueState(m.Sender.ID, TuesdayRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter the same way for tuesday",
				),
			)
		case TuesdayRequest:
			newTimetable[time.Tuesday] = utils.ParseSubjects(m.Text)

			err := db.SetDialogueState(m.Sender.ID, WednesdayRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter the same way for wednesday",
				),
			)
		case WednesdayRequest:
			newTimetable[time.Wednesday] = utils.ParseSubjects(m.Text)

			err := db.SetDialogueState(m.Sender.ID, ThursdayRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter the same way for Thursday",
				),
			)
		case ThursdayRequest:
			newTimetable[time.Thursday] = utils.ParseSubjects(m.Text)

			err := db.SetDialogueState(m.Sender.ID, FridayRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter the same way for friday",
				),
			)
		case FridayRequest:
			newTimetable[time.Friday] = utils.ParseSubjects(m.Text)

			err := db.SetDialogueState(m.Sender.ID, SaturdayRequest)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"Enter the same way for saturday",
				),
			)
		case SaturdayRequest:
			newTimetable[time.Saturday] = utils.ParseSubjects(m.Text)

			err := db.SetDialogueState(m.Sender.ID, NoRequest)
			if err != nil {
				panic(err)
			}

			_, err = db.CreateTimetable(m.Sender.ID, createdSchedule["id"].(int64), newTimetable)
			if err != nil {
				panic(err)
			}

			handleSendError(
				bot.Send(
					m.Chat,
					"All got in",
				),
			)
		}
	})
}
