package models

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Student struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

func (Student) Create(user *tb.User, classID int64) (neo4j.Node, error) {
	var student neo4j.Node

	result, err := db.Session.Run(
		"MATCH (c:Class { id: $class_id })" +
		"\n" +
		"CREATE (s:Student {" +
			"id: $id," +
			"first_name: $first_name," +
			"last_name: $last_name," +
			"username: $username," +
			"language_code: $language_code" +
		"})-[:STUDYING_IN]->(c)" +
		"\n" +
		"RETURN s",
		map[string]interface{}{
			"id": user.ID,
			"first_name": user.FirstName,
			"last_name": user.LastName,
			"username": user.Username,
			"language_code": user.LanguageCode,
			"class_id": classID,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		student = result.Record().GetByIndex(0).(neo4j.Node)
	}

	return student, nil
}
