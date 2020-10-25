package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	keyboard = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	homeworkButton  = keyboard.Text("Homework")
	timesheetButton = keyboard.Text("Timesheet")
	alertButton     = keyboard.Text("Alert")
	settingsButton  = keyboard.Text("Settings")
)

func registerKeyboard() {
	keyboard.Reply(
		keyboard.Row(homeworkButton, timesheetButton),
		keyboard.Row(alertButton, settingsButton),
	)

	bot.Handle(&homeworkButton, homeworkButtonHandler)
	bot.Handle(&timesheetButton, timesheetButtonHandler)
	bot.Handle(&alertButton, alertButtonHandler)
	bot.Handle(&settingsButton, settingsButtonHandler)
}

func homeworkButtonHandler(m *tb.Message) {
	handleBotError(
		bot.Send(
			m.Chat,
			"Handled homework button",
			&tb.ReplyMarkup{
				InlineKeyboard: generateInlineKeyboard("432"),
			},
		),
	)
}

func timesheetButtonHandler(m *tb.Message) {
	handleBotError(
		bot.Send(
			m.Chat,
			"Handled timesheet button",
		),
	)
}

func alertButtonHandler(m *tb.Message) {
	handleBotError(
		bot.Send(
			m.Chat,
			"Handled alert button",
		),
	)
}

func settingsButtonHandler(m *tb.Message) {
	handleBotError(
		bot.Send(
			m.Chat,
			"Handled settings button",
		),
	)
}
