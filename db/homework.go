package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CreateHomework(data map[string]interface{}) (neo4j.Node, error) {
	var homework neo4j.Node

	result, err := Session.Run(
		"MATCH (s:Student { tg_id: $tg_id })--(c:Class)"+
			"\n"+
			"MERGE (s)-[:CREATED]->(h:Homework { subject: $subject, task: $task, day: $day})-[:BELONGS_TO]-(c)"+
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

func QueryHomework() (neo4j.Node, error) {
	var homeworks neo4j.Node

	result, err := Session.Run(
		"MATCH (s:Student)-[:HAS]->(h:Homework)"+
			"\n"+
			"RETURN *",
		map[string]interface{}{},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		homeworks = result.Record().GetByIndex(0).(neo4j.Node)
	}

	return homeworks, result.Err()
}
