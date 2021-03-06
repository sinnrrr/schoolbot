package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	messageAPI = "https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s"

	weekdayInlineButtonText = "%s (%d.%d)"
	weekdayInlineButtonData = `{"date":%d,"action":%d}`

	homeworkReply = "\n" +
		"*%d.* _%s_" +
		"\n" +
		"*%s:* %s" +
		"\n" +
		"*%s:* %d.%d.%d" +
		"\n"

	alertReply = "\n" +
		"*%d.* _%s_" +
		"\n" +
		"*%s:* %d.%d.%d %d:%d" +
		"\n"

	cronAlert = "*Alert*" +
		"\n" +
		"%s"
)

func GenerateCronAlert(alert string) string {
	return fmt.Sprintf(cronAlert, alert)
}

func GenerateMessageURL(botToken string, chatID int, message string) string {
	return fmt.Sprintf(messageAPI, botToken, chatID, url.QueryEscape(message))
}

func GenerateScheduleMessage(schedule map[string]interface{}) string {
	var (
		reply = ""
		sliceSchedule = schedule["data"].([]interface{})
	)

	for index, scheduleTime := range sliceSchedule{
		reply += scheduleTime.(string)

		if (index+1)%2 == 0 {
			reply += "\n"
		} else {
			reply += " - "
		}
	}

	return reply
}

func GenerateTimetableMessage(subjects map[string]interface{}) string {
	l.SetDomain("weekdays")

	var (
		reply = []string{
			"*" + l.Gettext(time.Monday.String()) + "*" + "\n",
			"*" + l.Gettext(time.Tuesday.String()) + "*" + "\n",
			"*" + l.Gettext(time.Wednesday.String()) + "*" + "\n",
			"*" + l.Gettext(time.Thursday.String()) + "*" + "\n",
			"*" + l.Gettext(time.Friday.String()) + "*" + "\n",
			"*" + l.Gettext(time.Saturday.String()) + "*" + "\n",
		}
	)

	for weekday, subjectInterfaces := range subjects {
		for _, subjectName := range subjectInterfaces.([]interface{}) {
			switch strings.Title(weekday) {
			case time.Monday.String():
				reply[0] += subjectName.(string) + "\n"
			case time.Tuesday.String():
				reply[1] += subjectName.(string) + "\n"
			case time.Wednesday.String():
				reply[2] += subjectName.(string) + "\n"
			case time.Thursday.String():
				reply[3] += subjectName.(string) + "\n"
			case time.Friday.String():
				reply[4] += subjectName.(string) + "\n"
			case time.Saturday.String():
				reply[5] += subjectName.(string) + "\n"
			}
		}
	}

	return strings.Join(reply, "\n")
}

func GenerateAlertMessage(alerts []map[string]interface{}) string {
	var reply = ""

	l.SetDomain("general")

	if alerts != nil {
		for index, alert := range alerts {
			alertDate := time.Unix(alert["date"].(int64), 0)

			reply += fmt.Sprintf(
				alertReply,
				index+1,
				alert["content"].(string),
				l.Gettext("Onto"),
				alertDate.Day(),
				alertDate.Month(),
				alertDate.Year(),
				alertDate.Hour(),
				alertDate.Minute(),
			)
		}

		return reply
	}

	return l.Gettext("No alerts detected")
}

func GenerateHomeworkMessage(homeworks []map[string]interface{}) string {
	var reply = ""

	l.SetDomain("general")

	if homeworks != nil {
		for index, homework := range homeworks {
			homeworkDeadline := time.Unix(homework["deadline"].(int64), 0)

			reply += fmt.Sprintf(
				homeworkReply,
				index+1,
				homework["subject"].(string),
				l.Gettext("Task"),
				homework["task"].(string),
				l.Gettext("Deadline"),
				homeworkDeadline.Day(),
				homeworkDeadline.Month(),
				homeworkDeadline.Year(),
			)
		}

		return reply
	}

	return l.Gettext("No homeworks detected")
}

func WeekdayInlineButtonText(
	monthName string,
	date time.Time,
) string {
	return fmt.Sprintf(
		weekdayInlineButtonText,
		monthName,
		date.Day(),
		date.Month(),
	)
}

func WeekdayInlineButtonData(
	date int64,
	action int64,
) string {
	return fmt.Sprintf(
		weekdayInlineButtonData,
		date,
		action,
	)
}
