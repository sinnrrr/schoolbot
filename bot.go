package main

import (
	"github.com/sinnrrr/schoolbot/handlers"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	menu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	getButton    = menu.Text("Get")
	addButton    = menu.Text("Add")
	updateButton = menu.Text("Update")
	deleteButton = menu.Text("Delete")
)

func RegisterMenu(bot *tb.Bot)  {
	menu.Reply(
		menu.Row(getButton, addButton),
		menu.Row(updateButton, deleteButton),
	)

	bot.Handle(&getButton, handlers.GetButton)
	bot.Handle(&addButton, handlers.AddButton)
	bot.Handle(&updateButton, handlers.UpdateButton)
	bot.Handle(&deleteButton, handlers.DeleteButton)
}