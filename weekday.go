package main

import (
	"encoding/json"
	"github.com/sinnrrr/schoolbot/db"
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
		Text:   l.Gettext("Back"),
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

	l.SetDomain("weekdays")
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

	mondayInlineButton.Text = WeekdayInlineButtonText(l.Gettext("Monday"), mondayDate)
	tuesdayInlineButton.Text = WeekdayInlineButtonText(l.Gettext("Tuesday"), tuesdayDate)
	wednesdayInlineButton.Text = WeekdayInlineButtonText(l.Gettext("Wednesday"), wednesdayDate)
	thursdayInlineButton.Text = WeekdayInlineButtonText(l.Gettext("Thursday"), thursdayDate)
	fridayInlineButton.Text = WeekdayInlineButtonText(l.Gettext("Friday"), fridayDate)
	saturdayInlineButton.Text = WeekdayInlineButtonText(l.Gettext("Saturday"), saturdayDate)

	mondayInlineButton.Data = WeekdayInlineButtonData(mondayDate.Unix(), actionInt)
	tuesdayInlineButton.Data = WeekdayInlineButtonData(tuesdayDate.Unix(), actionInt)
	wednesdayInlineButton.Data = WeekdayInlineButtonData(wednesdayDate.Unix(), actionInt)
	thursdayInlineButton.Data = WeekdayInlineButtonData(thursdayDate.Unix(), actionInt)
	fridayInlineButton.Data = WeekdayInlineButtonData(fridayDate.Unix(), actionInt)
	saturdayInlineButton.Data = WeekdayInlineButtonData(saturdayDate.Unix(), actionInt)
}

func generateWeekdayInlineKeyboard(action string) *tb.ReplyMarkup {
	l.SetDomain("weekday")
	defineWeekdayInlineButtons(action)

	return &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{
			{mondayInlineButton, thursdayInlineButton},
			{tuesdayInlineButton, fridayInlineButton},
			{wednesdayInlineButton, saturdayInlineButton},
		},
	}
}

func weekdayInlineButtonHandler(c *tb.Callback) {
	l.SetDomain("dialogue")

	if c.Data == "back" {
		handleSendError(
			bot.Edit(
				c.Message,
				l.Gettext("What do you want to create today, master?"),
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
					l.Gettext("Send me subject name"),
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
					string(
						l.DGetdata(
							"examples",
							"time_enter.txt",
						),
					),
				),
			)
		}
	}
}
