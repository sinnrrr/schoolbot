package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func QueryHomework() (neo4j.Node, error) {
	var homeworks neo4j.Node

	result, err := Session.Run(
	"MATCH (s:Student)-[:HAS]->(h:Homework)" +
		"\n" +
		"RETURN *",
		map[string]interface{}{},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		homeworks = result.Record().GetByIndex(0).(neo4j.Node)
	}

	return homeworks, nil
}