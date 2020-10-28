package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	keyboard = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	newButton       = keyboard.Text("New")
	homeworkButton  = keyboard.Text("Homework")
	timetableButton = keyboard.Text("Timetable")
	alertButton     = keyboard.Text("Alert")
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
			"Handled new button",
			&tb.ReplyMarkup{
				InlineKeyboard: operationInlineKeyboard,
			},
		),
	)
}

func homeworkButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled homework button",
			&tb.ReplyMarkup{
				InlineKeyboard: generateActionsInlineKeyboard(432),
			},
		),
	)
}

func timetableButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled timetable button",
		),
	)
}

func alertButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled alert button",
		),
	)
}

func settingsButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled settings button",
		),
	)
}
