package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	updateInlineButton = tb.InlineButton{
		Unique: "update",
		Text:   "Update",
	}

	deleteInlineButton = tb.InlineButton{
		Unique: "delete",
		Text:   "Delete",
	}
)

func registerInlineKeyboard() {
	bot.Handle(&updateInlineButton, updateInlineButtonHandler)
	bot.Handle(&deleteInlineButton, deleteInlineButtonHandler)
}

func updateInlineButtonHandler(c *tb.Callback) {
	fmt.Println(c.Data)

	handleBotError(
		bot.Send(
			c.Sender,
			"Handled update inline button with data:",
		),
	)
}

func deleteInlineButtonHandler(c *tb.Callback) {
	handleBotError(
		bot.Send(
			c.Sender,
			"Handled delete inline button with data:",
		),
	)
}

func generateInlineKeyboard(data string) [][]tb.InlineButton {
	updateInlineButton.Data, deleteInlineButton.Data = data, data

	return [][]tb.InlineButton{{
		updateInlineButton,
		deleteInlineButton,
	}}
}
