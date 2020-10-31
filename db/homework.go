package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CreateHomework(data map[string]interface{}) (neo4j.Node, error) {
	var homework neo4j.Node

	result, err := Session.Run(
		"MATCH (s:Student { tg_id: $tg_id })--(c:Class)"+
			"\n"+
			"MERGE (s)-[:CREATED]->(h:Homework { subject: $subject, task: $task, deadline: $deadline})-[:BELONGS_TO]-(c)"+
			"\n"+
			"RETURN h",
		data,
	)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		homework = result.Record().GetByIndex(0).(neo4j.Node)
	}

	return homework, result.Err()
}

func QueryHomework(studentID int) ([]interface{}, error) {
	var homeworks []interface{}

	result, err := Session.Run(
		"MATCH (:Student { tg_id: $tg_id })-[:STUDYING_IN]->(:Class)<-[:BELONGS_TO]-(h:Homework)"+
			"\n"+
			"RETURN h"+
			"\n"+
			"ORDER BY h.deadline",
		map[string]interface{}{
			"tg_id": studentID,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		homeworks = append(homeworks, result.Record().Values()[0])
	}

	return homeworks, result.Err()
}

func DeleteHomework(homeworkID int) error {
	result, err := Session.Run(
		"MATCH (h:Homework)"+
			"\n"+
			"WHERE ID(h)=$id"+
			"\n"+
			"DETACH DELETE h",
		map[string]interface{}{
			"id": homeworkID,
		},
	)
	if err != nil {
		return err
	}

	//TODO: refactor answer without update/delete and one message

	return result.Err()
}
