package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

const (
	homeworkAction = 1
	alertAction = 2
)

var (
	item struct {
		ID     int
		Action int8
		Date   int64
	}

	homeworkInlineButton = tb.InlineButton{
		Data:   strconv.Itoa(homeworkAction),
		Unique: "newHomework",
		Text:   "Homework",
	}

	alertInlineButton = tb.InlineButton{
		Data:   strconv.Itoa(alertAction),
		Unique: "newAlert",
		Text:   "Alert",
	}

	operationInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{homeworkInlineButton},
			{alertInlineButton},
		},
	}
)

func registerInlineKeyboard() {
	bot.Handle(&homeworkInlineButton, operationInlineButtonHandler)
	bot.Handle(&alertInlineButton, operationInlineButtonHandler)

	bot.Handle(&mondayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&tuesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&wednesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&thursdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&fridayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&saturdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&weekdayBackInlineButton, weekdayInlineButtonHandler)
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
			generateWeekdayInlineKeyboard(c.Data),
		),
	)
}
