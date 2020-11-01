package main

import (
	"encoding/json"
	"github.com/sinnrrr/schoolbot/db"
	"github.com/sinnrrr/schoolbot/templates"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

var (
	mondayInlineButton = tb.InlineButton{
		Unique: "monday",
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
)

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

	mondayInlineButton.Text = templates.WeekdayInlineButtonText("Monday", mondayDate)
	tuesdayInlineButton.Text = templates.WeekdayInlineButtonText("Tuesday", tuesdayDate)
	wednesdayInlineButton.Text = templates.WeekdayInlineButtonText("Wednesday", wednesdayDate)
	thursdayInlineButton.Text = templates.WeekdayInlineButtonText("Thursday", thursdayDate)
	fridayInlineButton.Text = templates.WeekdayInlineButtonText("Friday", fridayDate)
	saturdayInlineButton.Text = templates.WeekdayInlineButtonText("Saturday", saturdayDate)

	mondayInlineButton.Data = templates.WeekdayInlineButtonData(mondayDate.Unix(), actionInt)
	tuesdayInlineButton.Data = templates.WeekdayInlineButtonData(tuesdayDate.Unix(), actionInt)
	wednesdayInlineButton.Data = templates.WeekdayInlineButtonData(wednesdayDate.Unix(), actionInt)
	thursdayInlineButton.Data = templates.WeekdayInlineButtonData(thursdayDate.Unix(), actionInt)
	fridayInlineButton.Data = templates.WeekdayInlineButtonData(fridayDate.Unix(), actionInt)
	saturdayInlineButton.Data = templates.WeekdayInlineButtonData(saturdayDate.Unix(), actionInt)
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

		switch item.Action {
		case homeworkAction:
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
				),
			)
		case alertAction:
			newAlert["date"] = item.Date
			err = db.SetDialogueState(c.Sender.ID, TimeRequest)
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
					"Now, send me time in 24-hour format."+
						" Note, that it should be rounded to 10 in order to properly work." +
						" Example: 15:10 or 8:40",
				),
			)
		}
	}
}
