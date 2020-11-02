package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

const (
	homeworkAction     = 1
	alertAction        = 2
	createLessonAction = 3
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

	createLessonInlineButton = tb.InlineButton{
		Data:   strconv.Itoa(createLessonAction),
		Unique: "createLesson",
		Text:   "Create schedule",
	}

	createLessonInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{{createLessonInlineButton}},
	}

	operationInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{homeworkInlineButton},
			{alertInlineButton},
		},
	}
)

func registerInlineKeyboard() {
	bot.Handle(&createLessonInlineButton, createLessonInlineButtonHandler)

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

func createLessonInlineButtonHandler(c *tb.Callback) {
	err := db.SetDialogueState(c.Sender.ID, ScheduleRequest)
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
			"Enter the start and the end of subjects"+
				"\n"+
				"Example: "+
				"\n"+
				"\n"+
				"8:30 9:10"+
				"\n"+
				"9:20 9:50"+
				"\n"+
				"10:00 10:40"+
				"\n"+
				"10:20 11:00",
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
			generateWeekdayInlineKeyboard(c.Data),
		),
	)
}
