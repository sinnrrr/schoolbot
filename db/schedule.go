package db

import "github.com/neo4j/neo4j-go-driver/neo4j"

func CreateSchedule(studentID int, data []string) (map[string]interface{}, error) {
	var schedule map[string]interface{}

	result, err := Session.Run(
		"MATCH (s:Student { tg_id: $tg_id })-[:STUDYING_IN]->(c:Class)"+
			"\n"+
			"MERGE (c)-[:USES]->(l:Schedule { data: $schedule })" +
			"\n" +
			"RETURN l",
		map[string]interface{}{
			"tg_id": studentID,
			"schedule": data,
		},
	)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		scheduleNode := result.Record().GetByIndex(0).(neo4j.Node)

		schedule = scheduleNode.Props()
		schedule["id"] = scheduleNode.Id()
	}

	return schedule, result.Err()
}
