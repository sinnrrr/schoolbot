package main

import (
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

const (
	homeworkAction        = 1
	alertAction           = 2
	createTimetableAction = 3
)

var (
	item struct {
		ID     int
		Action int8
		Date   int64
	}

	languageInlineButton = tb.InlineButton{
		Unique: "language",
		Text: "Language",
	}

	ukrainianInlineButton = tb.InlineButton{
		Data:   "uk",
		Unique: "ukrainian",
		Text:   "Ukrainian",
	}

	englishInlineButton = tb.InlineButton{
		Data:   "en",
		Unique: "english",
		Text:   "English",
	}

	homeworkInlineButton = tb.InlineButton{
		Data:   strconv.Itoa(homeworkAction),
		Unique: "newHomework",
		Text:   l.Gettext("Homework"),
	}

	alertInlineButton = tb.InlineButton{
		Data:   strconv.Itoa(alertAction),
		Unique: "newAlert",
		Text:   l.Gettext("Alert"),
	}

	createTimetableInlineButton = tb.InlineButton{
		Data:   strconv.Itoa(createTimetableAction),
		Unique: "createTimetable",
		Text:   l.Gettext("Create timetable"),
	}

	createTimetableInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{{createTimetableInlineButton}},
	}

	operationInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{homeworkInlineButton},
			{alertInlineButton},
		},
	}

	settingsInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{languageInlineButton},
		},
	}

	languagesInlineKeyboard = &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{englishInlineButton},
			{ukrainianInlineButton},
		},
	}
)

func registerInlineKeyboard() {
	l.SetDomain("dialogue")

	bot.Handle(&createTimetableInlineButton, createTimetableInlineButtonHandler)
	bot.Handle(&homeworkInlineButton, operationInlineButtonHandler)
	bot.Handle(&alertInlineButton, operationInlineButtonHandler)
	bot.Handle(&mondayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&tuesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&wednesdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&thursdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&fridayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&saturdayInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&weekdayBackInlineButton, weekdayInlineButtonHandler)
	bot.Handle(&languageInlineButton, languageInlineButtonHandler)
	bot.Handle(&englishInlineButton, languageInlineButtonHandler)
	bot.Handle(&ukrainianInlineButton, languageInlineButtonHandler)
}

func createTimetableInlineButtonHandler(c *tb.Callback) {
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
			string(
				l.DGetdata(
					"examples",
					"schedule_enter.txt",
				),
			),
		),
	)
}

func languageInlineButtonHandler(c *tb.Callback) {
	l.SetLanguage(c.Data)

	handleSendError(
		bot.Edit(
			c.Message,
			l.Gettext("Language was successfully changed"),
			keyboard,
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
			l.Gettext("Choose the day of the week ;)"),
			generateWeekdayInlineKeyboard(c.Data),
		),
	)
}

//func generateLanguagesSlice() {
//	for _, languageTag := range supportedLocale {
//		languagesSlice[languageTag] = tb.InlineButton{
//			Unique: strings.ToLower(languageTag.String()),
//			Text: languageTag.String(),
//			Data: languageTag.String(),
//		}
//	}
//
//	for _, button := range languagesSlice {
//		bot.Handle(&button, languageInlineButtonHandler)
//	}
//}

//func languageInlineButtonHandler(c *tb.Callback) {
//	handleSendError(
//		bot.Send(
//			c.Sender,
//			"Languge",
//		),
//	)
//}
