package utils

import (
	"bufio"
	"strings"
)

func ParseSchedule(message string) []string {
	var schedule []string

	scanner := bufio.NewScanner(strings.NewReader(message))
	for scanner.Scan() {
		scheduleComponents := strings.Split(scanner.Text(), " ")
		schedule = append(schedule, scheduleComponents[0], scheduleComponents[1])
	}

	return schedule
}

func ParseSubjects(message string) []string {
	var subjects []string

	scanner := bufio.NewScanner(strings.NewReader(message))
	for scanner.Scan() {
		subjects = append(subjects, scanner.Text())
	}

	return subjects
}

//func GenerateSchedule(
//	timetable map[time.Weekday][]string,
//	schedule [][]string,
//) map[string]map[string][]string {
//	var generated = make(map[string]map[string][]string)
//
//	for weekday, subjects := range timetable{
//		for index, subject := range subjects{
//			generated[weekday.String()] = map[string][]string{
//				subject: schedule[index],
//			}
//		}
//	}
//
//	return generated
//}
