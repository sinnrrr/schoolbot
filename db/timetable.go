package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"time"
)

func CreateTimetable(studentID int, scheduleID int64, data map[time.Weekday][]string) (map[string]interface{}, error) {
	var schedule map[string]interface{}

	result, err := Session.Run(
		"MATCH (l:Schedule) WHERE ID(l)=$schedule_id" +
			"\n" +
			"MATCH (s:Student { tg_id: $tg_id })-[:STUDYING_IN]->(c:Class)"+
			"\n"+
			"MERGE (s)-[:CREATED]->(t:Timetable {"+
			"monday: $monday,"+
			"tuesday: $tuesday,"+
			"wednesday: $wednesday,"+
			"thursday: $thursday,"+
			"friday: $friday,"+
			"saturday: $saturday"+
			"})<-[:STUDIES_ON]->(c)"+
			"\n" +
			"MERGE (t)-[:IMPLEMENTS]->(l)" +
			"\n"+
			"RETURN t",
		map[string]interface{}{
			"tg_id":  studentID,
			"schedule_id": scheduleID,
			"monday": data[time.Monday],
			"tuesday": data[time.Tuesday],
			"wednesday": data[time.Wednesday],
			"thursday": data[time.Thursday],
			"friday": data[time.Friday],
			"saturday": data[time.Saturday],
		},
	)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		schedule = result.Record().GetByIndex(0).(neo4j.Node).Props()
	}

	return schedule, result.Err()
}
