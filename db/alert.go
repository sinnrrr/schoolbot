package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/utils"
)

func QueryAlert(studentID int) ([]map[string]interface{}, error) {
	var alerts []map[string]interface{}

	result, err := Session.Run(
		"MATCH (:Student { tg_id: $tg_id })-[:STUDYING_IN]->(:Class)<-[:BELONGS_TO]-(a:Alert)\n"+
			"RETURN a\n"+
			"ORDER BY a.date",
		map[string]interface{}{
			"tg_id": studentID,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		alerts = append(alerts, result.Record().Values()[0].(neo4j.Node).Props())
	}

	return alerts, result.Err()
}

func CreateAlert(
	studentID int,
	alert map[string]interface{},
) (map[string]interface{}, error) {
	var createdAlert map[string]interface{}

	alertTime, err := utils.ParseTime(alert)
	if err != nil {
		return nil, err
	}

	result, err := Session.Run(
		"MATCH (s:Student { tg_id: $tg_id })\n"+
			"MATCH (s)-[:STUDYING_IN]->(c:Class)\n"+
			"CREATE (s)-[:CREATED]->(a:Alert { date: $date, content: $content })-[:BELONGS_TO]->(c)\n"+
			"RETURN a",
		map[string]interface{}{
			"tg_id":   studentID,
			"date":    alertTime.Unix(),
			"content": alert["content"].(string),
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		createdAlert = result.Record().GetByIndex(0).(neo4j.Node).Props()
	}

	return createdAlert, result.Err()
}

