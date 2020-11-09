package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseTime(alert map[string]interface{}) (time.Time, error) {
	parsedTime := strings.Split(alert["time"].(string), ":")
	parsedHour, err := strconv.Atoi(parsedTime[0])
	parsedMinute, err := strconv.Atoi(parsedTime[1])

	err = validateTime(parsedHour, parsedMinute)
	if err != nil {
		return time.Time{}, err
	}

	parsedDate := time.Unix(alert["date"].(int64), 0)
	return time.Date(
		parsedDate.Year(),
		parsedDate.Month(),
		parsedDate.Day(),
		parsedHour,
		parsedMinute,
		0,
		0,
		parsedDate.Location(),
	), err
}

func validateTime(hour int, minute int) error {
	if hour <= 23 && minute < 60 {
		return nil
	}

	return fmt.Errorf("incorrect value")
}
