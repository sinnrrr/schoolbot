package handlers

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/db"
)

func QueryHomework() (neo4j.Node, error) {
	var homeworks neo4j.Node

	result, err := db.Session.Run(
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