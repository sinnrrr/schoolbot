package main

import (
	"encoding/json"
	"fmt"
	"github.com/sinnrrr/schoolbot/handlers"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

var (
	data struct {
		ID     int
		Action int8
		Day    int8
	}

	homeworkInlineButton = tb.InlineButton{
		Data:   "1",
		Unique: "homework",
		Text:   "Homework",
	}

	timetableInlineButton = tb.InlineButton{
		Data:   "2",
		Unique: "timetable",
		Text:   "Timetable",
	}

	alertInlineButton = tb.InlineButton{
		Data:   "3",
		Unique: "alert",
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

	//weekdayBackInlineButton = tb.InlineButton{
	//	Unique: "back",
	//	Text:   "Back",
	//}
)

func registerInlineKeyboard() {
	bot.Handle(&updateActionInlineButton, actionInlineButtonHandler)
	bot.Handle(&deleteActionInlineButton, actionInlineButtonHandler)

	bot.Handle(&homeworkInlineButton, operationInlineButtonHandler)
	bot.Handle(&timetableInlineButton, operationInlineButtonHandler)
	bot.Handle(&alertInlineButton, operationInlineButtonHandler)
	//bot.Handle(&weekdayBackInlineButton, operationInlineButtonHandler)

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
				strconv.Itoa(int(data.Action))+
				" with ID "+
				string(rune(data.ID)),
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

	homework["day"] = data.Day
	err = handlers.SetDialogueState(c.Sender.ID, SubjectRequest)
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
			"Send me subject name",
		),
	)
}

func generateWeekdayInlineKeyboard(action string) [][]tb.InlineButton {
	actionInt, err := strconv.ParseInt(action, 10, 8)
	if err != nil {
		panic(err)
	}

	mondayInlineButton.Data = fmt.Sprintf(`{"day":1,"action":%d}`, actionInt)
	tuesdayInlineButton.Data = fmt.Sprintf(`{"day":2,"action":%d}`, actionInt)
	wednesdayInlineButton.Data = fmt.Sprintf(`{"day":3,"action":%d}`, actionInt)
	thursdayInlineButton.Data = fmt.Sprintf(`{"day":4,"action":%d}`, actionInt)
	fridayInlineButton.Data = fmt.Sprintf(`{"day":5,"action":%d}`, actionInt)
	saturdayInlineButton.Data = fmt.Sprintf(`{"day":6,"action":%d}`, actionInt)

	//weekdayBackInlineButton.Data = action

	return [][]tb.InlineButton{
		{mondayInlineButton, tuesdayInlineButton},
		{wednesdayInlineButton, thursdayInlineButton},
		{fridayInlineButton, saturdayInlineButton},
		//{weekdayBackInlineButton},
	}
}

func generateActionsInlineKeyboard(id int) [][]tb.InlineButton {
	updateActionInlineButton.Data = fmt.Sprintf(`{"id":%d,"action":"update"}`, id)
	deleteActionInlineButton.Data = fmt.Sprintf(`{"id":%d,"action":"delete"}`, id)

	return [][]tb.InlineButton{{
		updateActionInlineButton,
		deleteActionInlineButton,
	}}
}
