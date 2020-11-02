package main

import (
	"github.com/sinnrrr/schoolbot/db"
	"github.com/sinnrrr/schoolbot/templates"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	keyboard = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	newButton       = keyboard.Text("New")
	homeworkButton  = keyboard.Text("Homeworks")
	timetableButton = keyboard.Text("Timetable")
	alertButton     = keyboard.Text("Alerts")
	settingsButton  = keyboard.Text("Settings")
)

func registerKeyboard() {
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
			"What do you want to create today, master?",
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
	handleSendError(
		bot.Send(
			m.Chat,
			"Any timetable has been found",
			createLessonInlineKeyboard,
		),
	)
	//if timetable == nil {
	//	handleSendError(
	//		bot.Send(
	//			m.Chat,
	//			"Any timetable has been found",
	//		),
	//	)
	//} else {
	//	handleSendError(
	//		bot.Send(
	//			m.Chat,
	//			"Handled timetable button",
	//		),
	//	)
	//}
}

func settingsButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled settings button",
		),
	)
}
