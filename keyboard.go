package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

var (
	keyboard = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	newButton       = keyboard.Text("New")
	homeworkButton  = keyboard.Text("Homework")
	timetableButton = keyboard.Text("Timetable")
	alertButton     = keyboard.Text("Alert")
	settingsButton  = keyboard.Text("Settings")
)

func registerKeyboard() {
	keyboard.Reply(
		keyboard.Row(newButton),
		keyboard.Row(homeworkButton, timetableButton),
		keyboard.Row(alertButton, settingsButton),
	)

	bot.Handle(&newButton, newButtonHandler)
	bot.Handle(&homeworkButton, homeworkButtonHandler)
	bot.Handle(&timetableButton, timetableButtonHandler)
	bot.Handle(&alertButton, alertButtonHandler)
	bot.Handle(&settingsButton, settingsButtonHandler)
}

func newButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"What do you want to create today, master?",
			operationInlineKeyboard,
		),
	)
}

func homeworkButtonHandler(m *tb.Message) {
	homeworks, err := db.QueryHomework(m.Sender.ID)
	if err != nil {
		panic(err)
	}

	if homeworks == nil {
		handleSendError(
			bot.Send(
				m.Chat,
				"No homeworks detected",
			),
		)
	} else {
		for _, homework := range homeworks {
			currentHomework := homework.(neo4j.Node).Props()
			currentHomeworkDeadline := time.Unix(currentHomework["deadline"].(int64), 0)

			handleSendError(
				bot.Send(
					m.Chat,
					"*Subject: *"+
						currentHomework["subject"].(string)+
						"\n"+
						"*Task: *"+
						currentHomework["task"].(string)+
						"\n"+
						"*Deadline: *"+
						fmt.Sprintf(
							"%d\\.%d\\.%d",
							currentHomeworkDeadline.Day(),
							currentHomeworkDeadline.Month(),
							currentHomeworkDeadline.Year(),
						),
					generateActionsInlineKeyboard(homework.(neo4j.Node).Id()),
					tb.ModeMarkdownV2,
				),
			)
		}
	}

}

func timetableButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled timetable button",
		),
	)
}

func alertButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Chat,
			"Handled alert button",
		),
	)
}

func settingsButtonHandler(m *tb.Message) {
	handleSendError(
		bot.Send(
			m.Sender,
			"Handled settings button",
		),
	)
}
