package main

import (
	"encoding/json"
	"fmt"
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

const (
	UpdateAction = 1
	DeleteAction = 2
)

var (
	item struct {
		ID     int
		Action int8
		Date   int64
	}

	homeworkInlineButton = tb.InlineButton{
		Data:   "1",
		Unique: "newHomework",
		Text:   "Homework",
	}

	alertInlineButton = tb.InlineButton{
		Data:   "2",
		Unique: "alert",
		Text:   "Alert",
	}

	cancelInlineButton = tb.InlineButton{
		Data:   "3",
		Unique: "cancel",
		Text:   "Cancel",
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
	}

	wednesdayInlineButton = tb.InlineButton{
		Unique: "wednesday",
	}

	thursdayInlineButton = tb.InlineButton{
		Unique: "thursday",
	}

	fridayInlineButton = tb.InlineButton{
		Unique: "friday",
	}

	saturdayInlineButton = tb.InlineButton{
		Unique: "saturday",
	}

	weekdayBackInlineButton = tb.InlineButton{
		Unique: "weekday_back",
		Data:   "back",
		Text:   "Back",
	}

	operationInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{homeworkInlineButton, alertInlineButton},
			{cancelInlineButton},
		},
	}
)

func registerInlineKeyboard() {
	bot.Handle(&updateActionInlineButton, actionInlineButtonHandler)
	bot.Handle(&deleteActionInlineButton, actionInlineButtonHandler)

	bot.Handle(&homeworkInlineButton, operationInlineButtonHandler)
	bot.Handle(&alertInlineButton, operationInlineButtonHandler)
	bot.Handle(&cancelInlineButton, operationInlineButtonHandler)

	bot.Handle(&mondayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&tuesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&wednesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&thursdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&fridayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&saturdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&weekdayBackInlineButton, weekdayInlineButtonHandler)
}

func actionInlineButtonHandler(c *tb.Callback) {
	err := json.Unmarshal([]byte(c.Data), &item)
	if err != nil {
		panic(err)
	}

	err = bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	if err != nil {
		panic(err)
	}

	switch item.Action {
	case UpdateAction:
		handleSendError(
			bot.Send(
				c.Sender,
				"Handled update action for homework with ID "+
					strconv.Itoa(item.ID),
			),
		)
	case DeleteAction:
		err := db.DeleteHomework(item.ID)
		if err != nil {
			panic(err)
		}

		handleSendError(
			bot.Edit(
				c.Message,
				"This homework was deleted",
			),
		)
	}
}

func operationInlineButtonHandler(c *tb.Callback) {
	err := bot.Respond(c, &tb.CallbackResponse{
		ShowAlert: false,
	})
	if err != nil {
		panic(err)
	}

	if c.Data == "3" {
		cancelOperation(c)
	} else {
		handleSendError(
			bot.Edit(
				c.Message,
				"Choose the day of the week",
				generateWeekdayInlineKeyboard(c.Data),
			),
		)
	}
}

func weekdayInlineButtonHandler(c *tb.Callback) {
	if c.Data == "back" {
		handleSendError(
			bot.Edit(
				c.Message,
				"What do you want to create today, master?",
				operationInlineKeyboard,
			),
		)
	} else {
		err := json.Unmarshal([]byte(c.Data), &item)
		if err != nil {
			panic(err)
		}

		newHomework["deadline"] = item.Date
		err = db.SetDialogueState(c.Sender.ID, SubjectRequest)
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
				&tb.ReplyMarkup{
					InlineKeyboard: [][]tb.InlineButton{{cancelInlineButton}},
				},
			),
		)
	}
}

func cancelOperation(c *tb.Callback) {
	err := bot.Delete(c.Message)
	if err != nil {
		panic(err)
	}

	handleSendError(bot.Send(c.Sender, "Cancelled operation", keyboard))
}

func generateWeekdayInlineKeyboard(action string) *tb.ReplyMarkup {
	defineWeekdayInlineButtons(action)

	return &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{mondayInlineButton, tuesdayInlineButton},
			{wednesdayInlineButton, thursdayInlineButton},
			{fridayInlineButton, saturdayInlineButton},
			{weekdayBackInlineButton},
		},
	}
}

func generateActionsInlineKeyboard(id int64) *tb.ReplyMarkup {
	updateActionInlineButton.Data = fmt.Sprintf(`{"id":%d,"action":1}`, id)
	deleteActionInlineButton.Data = fmt.Sprintf(`{"id":%d,"action":2}`, id)

	return &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{{
			updateActionInlineButton,
			deleteActionInlineButton,
		}},
	}
}

func defineWeekdayInlineButtons(action string) {
	var (
		mondayDate    time.Time
		tuesdayDate   time.Time
		wednesdayDate time.Time
		thursdayDate  time.Time
		fridayDate    time.Time
		saturdayDate  time.Time
	)

	currentDate := time.Now()

	actionInt, err := strconv.ParseInt(action, 10, 8)
	if err != nil {
		panic(err)
	}

	switch time.Now().Weekday() {
	case time.Monday:
		mondayDate = currentDate
		tuesdayDate = mondayDate.AddDate(0, 0, 1)
		wednesdayDate = tuesdayDate.AddDate(0, 0, 1)
		thursdayDate = wednesdayDate.AddDate(0, 0, 1)
		fridayDate = thursdayDate.AddDate(0, 0, 1)
		saturdayDate = fridayDate.AddDate(0, 0, 1)
	case time.Tuesday:
		tuesdayDate = currentDate
		wednesdayDate = tuesdayDate.AddDate(0, 0, 1)
		thursdayDate = wednesdayDate.AddDate(0, 0, 1)
		fridayDate = thursdayDate.AddDate(0, 0, 1)
		saturdayDate = fridayDate.AddDate(0, 0, 1)
		mondayDate = saturdayDate.AddDate(0, 0, 2)
	case time.Wednesday:
		wednesdayDate = currentDate
		thursdayDate = wednesdayDate.AddDate(0, 0, 1)
		fridayDate = thursdayDate.AddDate(0, 0, 1)
		saturdayDate = fridayDate.AddDate(0, 0, 1)
		mondayDate = saturdayDate.AddDate(0, 0, 2)
		tuesdayDate = mondayDate.AddDate(0, 0, 1)
	case time.Thursday:
		thursdayDate = currentDate
		fridayDate = thursdayDate.AddDate(0, 0, 1)
		saturdayDate = fridayDate.AddDate(0, 0, 1)
		mondayDate = saturdayDate.AddDate(0, 0, 2)
		tuesdayDate = mondayDate.AddDate(0, 0, 1)
		wednesdayDate = tuesdayDate.AddDate(0, 0, 1)
	case time.Friday:
		fridayDate = currentDate
		saturdayDate = fridayDate.AddDate(0, 0, 1)
		mondayDate = saturdayDate.AddDate(0, 0, 2)
		tuesdayDate = mondayDate.AddDate(0, 0, 1)
		wednesdayDate = tuesdayDate.AddDate(0, 0, 1)
		thursdayDate = wednesdayDate.AddDate(0, 0, 1)
	case time.Saturday:
		saturdayDate = currentDate
		mondayDate = saturdayDate.AddDate(0, 0, 2)
		tuesdayDate = mondayDate.AddDate(0, 0, 1)
		wednesdayDate = tuesdayDate.AddDate(0, 0, 1)
		thursdayDate = wednesdayDate.AddDate(0, 0, 1)
		fridayDate = thursdayDate.AddDate(0, 0, 1)
	}

	mondayInlineButton.Text = fmt.Sprintf("Monday (%d.%d)", mondayDate.Day(), mondayDate.Month())
	tuesdayInlineButton.Text = fmt.Sprintf("Tuesday (%d.%d)", tuesdayDate.Day(), tuesdayDate.Month())
	wednesdayInlineButton.Text = fmt.Sprintf("Wednesday (%d.%d)", wednesdayDate.Day(), wednesdayDate.Month())
	thursdayInlineButton.Text = fmt.Sprintf("Thursday (%d.%d)", thursdayDate.Day(), thursdayDate.Month())
	fridayInlineButton.Text = fmt.Sprintf("Friday (%d.%d)", fridayDate.Day(), fridayDate.Month())
	saturdayInlineButton.Text = fmt.Sprintf("Saturday (%d.%d)", saturdayDate.Day(), saturdayDate.Month())

	mondayInlineButton.Data = fmt.Sprintf(`{"date":%d,"action":%d}`, mondayDate.Unix(), actionInt)
	tuesdayInlineButton.Data = fmt.Sprintf(`{"date":%d,"action":%d}`, tuesdayDate.Unix(), actionInt)
	wednesdayInlineButton.Data = fmt.Sprintf(`{"date":%d,"action":%d}`, wednesdayDate.Unix(), actionInt)
	thursdayInlineButton.Data = fmt.Sprintf(`{"date":%d,"action":%d}`, thursdayDate.Unix(), actionInt)
	fridayInlineButton.Data = fmt.Sprintf(`{"date":%d,"action":%d}`, fridayDate.Unix(), actionInt)
	saturdayInlineButton.Data = fmt.Sprintf(`{"date":%d,"action":%d}`, saturdayDate.Unix(), actionInt)
}
