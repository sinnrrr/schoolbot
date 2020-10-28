package main

import (
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	data struct {
		ID     string
		Name   string
		Action string
	}

	homeworkInlineButton = tb.InlineButton{
		Unique: "homework",
		Data:   "homework",
		Text:   "Homework",
	}

	timetableInlineButton = tb.InlineButton{
		Unique: "timetable",
		Data:   "timetable",
		Text:   "Timetable",
	}

	alertInlineButton = tb.InlineButton{
		Unique: "alert",
		Data:   "alert",
		Text:   "Alert",
	}

	updateActionInlineButton = tb.InlineButton{
		Unique: "update",
		Text:   "Update",
	}

	deleteActionInlineButton = tb.InlineButton{
		Unique: "delete",
		Text:   "Delete",
	}

	mondayInlineButton = tb.InlineButton{
		Unique: "monday",
		Text:   "Monday",
	}

	tuesdayInlineButton = tb.InlineButton{
		Unique: "tuesday",
		Text:   "Tuesday",
	}

	wednesdayInlineButton = tb.InlineButton{
		Unique: "wednesday",
		Text:   "Wednesday",
	}

	thursdayInlineButton = tb.InlineButton{
		Unique: "thursday",
		Text:   "Thursday",
	}

	fridayInlineButton = tb.InlineButton{
		Unique: "friday",
		Text:   "Friday",
	}

	saturdayInlineButton = tb.InlineButton{
		Unique: "saturday",
		Text:   "Saturday",
	}

	operationInlineKeyboard = [][]tb.InlineButton{
		{homeworkInlineButton},
		{timetableInlineButton},
		{alertInlineButton},
	}

	weekdayBackInlineButton = tb.InlineButton{
		Unique: "back",
		Text:   "Back",
	}
)

func registerInlineKeyboard() {
	bot.Handle(&updateActionInlineButton, actionInlineButtonHandler)
	bot.Handle(&deleteActionInlineButton, actionInlineButtonHandler)

	bot.Handle(&homeworkInlineButton, operationInlineButtonHandler)
	bot.Handle(&timetableInlineButton, operationInlineButtonHandler)
	bot.Handle(&alertInlineButton, operationInlineButtonHandler)
	bot.Handle(&weekdayBackInlineButton, operationInlineButtonHandler)

	bot.Handle(&mondayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&tuesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&wednesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&thursdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&fridayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&saturdayInlineButton, weekdayInlineButtonHandler)
}

func actionInlineButtonHandler(c *tb.Callback) {
	err := json.Unmarshal([]byte(c.Data), &data)
	if err != nil {
		panic(err)
	}

	err = bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	if err != nil {
		panic(err)
	}

	handleSendError(
		bot.Send(
			c.Sender,
			"Handled action "+
				data.Action+
				" with ID "+
				data.ID,
		),
	)
}

func operationInlineButtonHandler(c *tb.Callback) {
	err := bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	if err != nil {
		panic(err)
	}

	handleSendError(
		bot.Edit(
			c.Message,
			"Choose the day of the week",
			&tb.ReplyMarkup{
				InlineKeyboard: generateWeekdayInlineKeyboard(c.Data),
			},
		),
	)
}

func weekdayInlineButtonHandler(c *tb.Callback) {
	err := json.Unmarshal([]byte(c.Data), &data)
	if err != nil {
		panic(err)
	}

	err = bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	if err != nil {
		panic(err)
	}

	handleSendError(
		bot.Edit(
			c.Message,
			"Handled " + data.Name + " inline button with action " + data.Action,
		),
	)
}

func generateWeekdayInlineKeyboard(action string) [][]tb.InlineButton {
	mondayInlineButton.Data = fmt.Sprintf(`{"name":"monday","action":"%s"}`, action)
	tuesdayInlineButton.Data = fmt.Sprintf(`{"name":"tuesday","action":"%s"}`, action)
	wednesdayInlineButton.Data = fmt.Sprintf(`{"name":"wednesday","action":"%s"}`, action)
	thursdayInlineButton.Data = fmt.Sprintf(`{"name":"thursday","action":"%s"}`, action)
	fridayInlineButton.Data = fmt.Sprintf(`{"name":"friday","action":"%s"}`, action)
	saturdayInlineButton.Data = fmt.Sprintf(`{"name":"saturday","action":"%s"}`, action)

	weekdayBackInlineButton.Data = action

	return [][]tb.InlineButton{
		{mondayInlineButton, tuesdayInlineButton},
		{wednesdayInlineButton, thursdayInlineButton},
		{fridayInlineButton, saturdayInlineButton},
		{weekdayBackInlineButton},
	}
}

func generateActionsInlineKeyboard(id string) [][]tb.InlineButton {
	updateActionInlineButton.Data = fmt.Sprintf(`{"id":"%s","action":"update"}`, id)
	deleteActionInlineButton.Data = fmt.Sprintf(`{"id":"%s","action":"delete"}`, id)

	return [][]tb.InlineButton{{
		updateActionInlineButton,
		deleteActionInlineButton,
	}}
}
