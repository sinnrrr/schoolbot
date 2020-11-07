package main

import (
	"github.com/sinnrrr/schoolbot/db"
	"github.com/sinnrrr/schoolbot/templates"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	keyboard = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	newButton       = keyboard.Text(l.Gettext("New"))
	homeworkButton  = keyboard.Text(l.Gettext("Homeworks"))
	timetableButton = keyboard.Text(l.Gettext("Timetable"))
	alertButton     = keyboard.Text(l.Gettext("Alerts"))
	settingsButton  = keyboard.Text(l.Gettext("Settings"))
)

func registerKeyboard() {
	l.SetDomain("dialogue")

	keyboard.Reply(
		keyboard.Row(newButton),
		keyboard.Row(homeworkButton, timetableButton),
		keyboard.Row(alertButton, settingsButton),
	)

	bot.Handle(&newButton, newButtonHandler)
	bot.Handle(&homeworkButton, homeworkButtonHandler)
	bot.Handle(&timetableButton, timetableButtonHandler)
	bot.Handle(&alertButton, alertButtonHandler)
	bot.Handle(&settingsButton, settingsButtonHandler)
}

func newButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			l.Gettext("What do you want to create today, master?"),
			operationInlineKeyboard,
		),
	)
}

func homeworkButtonHandler(m *tb.Message) {
	homeworks, err := db.QueryHomework(m.Sender.ID)
	if err != nil {
		panic(err)
	}

	handleSendError(
		bot.Send(
			m.Chat,
			templates.GenerateHomeworkMessage(homeworks),
			tb.ModeMarkdown,
		),
	)
}

func alertButtonHandler(m *tb.Message) {
	alerts, err := db.QueryAlert(m.Sender.ID)
	if err != nil {
		panic(err)
	}

	handleSendError(
		bot.Send(
			m.Chat,
			templates.GenerateAlertMessage(alerts),
			tb.ModeMarkdown,
		),
	)
}

func timetableButtonHandler(m *tb.Message) {
	timetable, err := db.StudentTimetable(m.Sender.ID)
	if err != nil {
		panic(err)
	}

	if timetable == nil {
		handleSendError(
			bot.Send(
				m.Chat,
				l.Gettext("Timetable for your class hasn't been created yet"),
				createTimetableInlineKeyboard,
			),
		)
	} else {
		handleSendError(
			bot.Send(
				m.Chat,
				templates.GenerateTimetableMessage(timetable[0]),
				tb.ModeMarkdown,
			),
		)
		handleSendError(
			bot.Send(
				m.Chat,
				templates.GenerateScheduleMessage(timetable[1]),
			),
		)
	}
}

func settingsButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			l.DGettext("general", "Coming soon :)"),
		),
	)
}
