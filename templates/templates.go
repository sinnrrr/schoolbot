package templates

import (
	"fmt"
	"time"
)

const (
	weekdayInlineButtonText = "%s (%d.%d)"
	weekdayInlineButtonData = `{"date":%d,"action":%d}`

	actionInlineButtonData = `{"id":%d,"action":%d}`

	homeworkReply = "\n" +
		"*%d.* _%s_" +
		"\n" +
		"*Task:* %s" +
		"\n" +
		"*Deadline:* %d.%d.%d" +
		"\n"

	alertReply = "\n" +
		"*%d.* _%s_" +
		"\n" +
		"*Onto:* %d.%d.%d" +
		"\n"
)

func GenerateAlertMessage(alerts []map[string]interface{}) string {
	var reply = ""

	if alerts != nil {
		for index, alert := range alerts {
			alertDate := time.Unix(alert["date"].(int64), 0)

			reply += fmt.Sprintf(
				alertReply,
				index+1,
				alert["content"].(string),
				alertDate.Day(),
				alertDate.Month(),
				alertDate.Year(),
			)
		}

		return reply
	}

	return "No alerts detected"
}

func GenerateHomeworkMessage(homeworks []map[string]interface{}) string {
	var reply = ""

	if homeworks != nil {
		for index, homework := range homeworks {
			homeworkDeadline := time.Unix(homework["deadline"].(int64), 0)

			reply += fmt.Sprintf(
				homeworkReply,
				index+1,
				homework["subject"].(string),
				homework["task"].(string),
				homeworkDeadline.Day(),
				homeworkDeadline.Month(),
				homeworkDeadline.Year(),
			)
		}

		return reply
	}

	return "No homeworks detected"
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

func ActionInlineButtonData(
	id int64,
	action int,
) string {
	return fmt.Sprintf(
		actionInlineButtonData,
		id,
		action,
	)
}
