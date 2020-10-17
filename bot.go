package main

import tb "gopkg.in/tucnak/telebot.v2"

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

	bot.Handle(&getButton, getButtonHandler)
	bot.Handle(&addButton, addButtonHandler)
	bot.Handle(&updateButton, updateButtonHandler)
	bot.Handle(&deleteButton, deleteButtonHandler)
}

func getButtonHandler(m *tb.Message) {
}

func addButtonHandler(m *tb.Message) {

}

func updateButtonHandler(m *tb.Message) {

}

func deleteButtonHandler(m *tb.Message) {

}