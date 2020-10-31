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
		"*Deadline:* %s" +
		"\n"
)

func GenerateHomeworkMessage(
	index int,
	deadline time.Time,
	subject string,
	task string,
) string {
	formattedDeadline := fmt.Sprintf(
		"%d.%d.%d",
		deadline.Day(),
		deadline.Month(),
		deadline.Year(),
	)

	return fmt.Sprintf(
		homeworkReply,
		index+1,
		subject,
		task,
		formattedDeadline,
	)
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
