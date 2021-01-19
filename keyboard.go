package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	keyboard         = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true}
)

func registerKeyboard() {
	l.SetDomain("general")

	var (
		newButton       = keyboard.Text(l.Gettext("New"))
		homeworkButton  = keyboard.Text(l.Gettext("Homeworks"))
		timetableButton = keyboard.Text(l.Gettext("Timetable"))
		alertButton     = keyboard.Text(l.Gettext("Alerts"))
		settingsButton  = keyboard.Text(l.Gettext("Settings"))
	)

	keyboard.Reply(
		keyboard.Row(newButton),
		keyboard.Row(homeworkButton, timetableButton),
		keyboard.Row(alertButton, settingsButton),
	)

	l.SetDomain("dialogue")

	bot.Handle(&newButton, newButtonHandler)
	bot.Handle(&homeworkButton, homeworkButtonHandler)
	bot.Handle(&timetableButton, timetableButtonHandler)
	bot.Handle(&alertButton, alertButtonHandler)
	bot.Handle(&settingsButton, settingsButtonHandler)
}

func newButtonHandler(m *tb.Message) {
	l.SetDomain("dialogue")

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
			GenerateHomeworkMessage(homeworks),
			keyboard,
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
			GenerateAlertMessage(alerts),
			keyboard,
			tb.ModeMarkdown,
		),
	)
}

func timetableButtonHandler(m *tb.Message) {
	l.SetDomain("dialogue")

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
				GenerateTimetableMessage(timetable[0]),
				tb.ModeMarkdown,
			),
		)
		handleSendError(
			bot.Send(
				m.Chat,
				GenerateScheduleMessage(timetable[1]),
				keyboard,
			),
		)
	}
}

func settingsButtonHandler(m *tb.Message) {
	l.SetDomain("general")

	handleSendError(
		bot.Send(
			m.Chat,
			l.Gettext("Go ahead, tweak me"),
			settingsInlineKeyboard,
		),
	)
}